/**
 * This code is the retranscription in typescript based on https://github.com/aerik/aerik.github.io
 */

import { Ref, ShallowRef } from 'vue';
import { Note, NotePlayed, notes } from './notes';
import { AddNotesDocument } from '@/gql/mutations/AddNote';
import { ApolloClient } from '@apollo/client/core/ApolloClient';
import { NormalizedCacheObject } from '@apollo/client/cache/inmemory/types';
import { StartPerformanceDocument } from '@/gql/mutations/StartPerformance';
import { FinishPerformanceDocument } from '@/gql/mutations/FinishPerformance';

let analyser: AnalyserNode;

let gainNode: GainNode;
let hertzBinSize: number;
let frequencyData: Uint8Array;
let buflen: number;

let lastNotePlayed: NotePlayed = { octave: 0, value: '', at: 0, duration: 0 };
let maxBytePlayed = 0; // max gain of note played
// canvas
const maxValue = 256; //based on Uint8Array possible values

/**
 * Initialise settings of the microphone
 * @returns Audio stream
 */
export const initMicrophone = async (
    songId: string,
    sensitivity: number,
    apolloClient: ApolloClient<NormalizedCacheObject>,
    perfID: Ref<string>
): Promise<MediaStream | null> => {
    const audioCtx: AudioContext = new window.AudioContext();

    analyser = audioCtx.createAnalyser();
    analyser.smoothingTimeConstant = 0.8; //default is 0.8, less is more responsive
    analyser.minDecibels = -sensitivity; //-100 is default and is more sensitive (more noise)
    analyser.fftSize = 8192 * 4; //need at least 8192 to detect differences in low notes

    const sampleRate: number = audioCtx.sampleRate;
    gainNode = audioCtx.createGain();
    gainNode.connect(audioCtx.destination);

    hertzBinSize = sampleRate / analyser.fftSize;
    frequencyData = new Uint8Array(analyser.frequencyBinCount);
    buflen = frequencyData.length;

    // ask for microphone permission
    let stream: MediaStream | null = await navigator.mediaDevices.getUserMedia({
        audio: {
            noiseSuppression: true,
            echoCancellation: true,
        },
    });

    const micSource: MediaStreamAudioSourceNode = audioCtx.createMediaStreamSource(stream);
    micSource.connect(gainNode);
    micSource.connect(analyser);

    // reset timestamps
    notes.forEach((note) => {
        note.timestamps = [];
    });

    // Start performance
    try {
        const result = await apolloClient.mutate({
            mutation: StartPerformanceDocument,
            variables: {
                songId: songId,
            },
        });
        if (result.data?.startPerformance.id) perfID.value = result.data?.startPerformance.id;
        console.log(perfID.value);
    } catch (e) {
        console.error(e);
        stream.getAudioTracks().forEach((track: MediaStreamTrack) => track.stop());
        stream = null;
    }

    return stream;
};

/**
 * Record played notes
 * @param startTimeStamp    Timestamp of the start of the track
 * @param enableCanvas      Enable canvans
 * @param decibelMin        Minimum binary value of decibel to track (between 0 and 255)
 */
export const getTones = (startTimeStamp: number, enableCanvas = false, decibelMin: number) => {
    analyser.getByteFrequencyData(frequencyData);
    let count = 0;
    let total = 0;
    const cutoff = 20; //redundant with decibels?
    let nPtr = 0; //notePointer
    for (let i = 0; i < buflen; i++) {
        const fdat: number = frequencyData[i];
        const freq: number = i * hertzBinSize; //freq in hertz for this sample
        const curNote: Note = notes[nPtr];
        const nextNote: Note = notes[nPtr + 1];
        //cut off halfway into the next note
        const hzLimit: number = curNote.frequency + (nextNote.frequency - curNote.frequency) / 2;
        if (freq < hzLimit) {
            if (fdat > cutoff) {
                count++;
                total += fdat;
            }
        } else {
            if (count > 0) {
                const power = total / count - notes[nPtr].noise;
                notes[nPtr].power = power > 0 ? power : 0; // this check is only done to not display aberrations in canvas
                // if a note is detected
                if (notes[nPtr].power > maxBytePlayed * 0.8 && notes[nPtr].power > decibelMin) {
                    maxBytePlayed = notes[nPtr].power;
                    const currentTimestamp = Date.now() - startTimeStamp;

                    // this condition aims to not log too much notes
                    if (
                        // if note is different
                        notes[nPtr].octave !== lastNotePlayed.octave ||
                        notes[nPtr].step !== lastNotePlayed.value ||
                        // if the same note has been played after 500 ms
                        (notes[nPtr].octave === lastNotePlayed.octave &&
                            notes[nPtr].step === lastNotePlayed.value &&
                            currentTimestamp - 500 >= lastNotePlayed.at)
                    ) {
                        /**
                         * Push note into an array and is processed later
                         * @see getDuration
                         */
                        notes[nPtr].timestamps.push(currentTimestamp);

                        //console.log(JSON.stringify(notes[nPtr]) + ' timestamp: ' + currentTimestamp);
                    }
                    // save the note that has been detected
                    lastNotePlayed = {
                        octave: notes[nPtr].octave,
                        value: notes[nPtr].step,
                        at: currentTimestamp,
                        duration: 0,
                    };
                }

                count = 0;
                total = 0;
            } else {
                notes[nPtr].power = 0;
            }
            //next note
            if (nPtr < notes.length - 2) {
                nPtr++;
            }
        }
    }
    if (enableCanvas) {
        const canvas: HTMLCanvasElement = <HTMLCanvasElement>document.getElementById('visualizer');
        const ctx: CanvasRenderingContext2D | null = canvas.getContext('2d');
        if (ctx === null) {
            throw new Error('Context of canvas not available');
        }
        canvas.height = notes.length * 10;
        canvas.width = maxValue + 120;
        ctx.textAlign = 'left';
        //ctx.clearRect(0,0,canvas.width, canvas.height);
        for (let n = 0; n < notes.length; n++) {
            ctx.fillText(notes[n].octave + ' ' + notes[n].step, 65, n * 10);
            ctx.save();
            ctx.textAlign = 'right';
            ctx.fillText(notes[n].frequency + ' Hz', 60, n * 10);
            ctx.restore();
        }
        ctx.save();
        //horizontal bars
        for (let n = 0; n < notes.length; n++) {
            const colString = 'hsl(' + (360 * n) / notes.length + ',100%,80%)';
            ctx.fillStyle = colString;
            ctx.fillRect(120, n * 10, notes[n].power, -10);
        }
        ctx.restore();
    }
};

/**
 * Record noise for phase opposition purpose
 */
export const getNoise = () => {
    analyser.getByteFrequencyData(frequencyData);
    let count = 0;
    let total = 0;
    const cutoff = 20; //redundant with decibels?
    let nPtr = 0; //notePointer
    for (let i = 0; i < buflen; i++) {
        const fdat: number = frequencyData[i];
        const freq: number = i * hertzBinSize; //freq in hertz for this sample
        const curNote: Note = notes[nPtr];
        const nextNote: Note = notes[nPtr + 1];
        //cut off halfway into the next note
        const hzLimit: number = curNote.frequency + (nextNote.frequency - curNote.frequency) / 2;
        if (freq < hzLimit) {
            if (fdat > cutoff) {
                count++;
                total += fdat;
            }
        } else {
            if (count > 0) {
                if (total / count > notes[nPtr].noise) {
                    notes[nPtr].noise = total / count;
                }
                count = 0;
                total = 0;
            }
            //next note
            if (nPtr < notes.length - 2) {
                nPtr++;
            }
        }
    }
};

/**
 * Get notes duration and send them to backend
 * @param stream Shallow reference of the audio stream
 * @param apolloClient Web socket client
 * @param perfID Performance ID
 * @param startTimeStamp timestamp when noise recording is finished
 * @param acquisitionDelay Sending to backend delay
 */
export const getDuration = (
    stream: ShallowRef<MediaStream | null>,
    apolloClient: ApolloClient<NormalizedCacheObject>,
    perfID: Ref<string>,
    startTimeStamp: number,
    acquisitionDelay: Ref<number>
) => {
    const currentTimestamp = Date.now() - startTimeStamp;
    const notesDetected: NotePlayed[] = [];

    notes.forEach((note) => {
        const timestamps = note.timestamps.filter((timestamp) => timestamp > currentTimestamp - acquisitionDelay.value);

        if (timestamps.length > 0) {
            const timestamp = timestamps[0];
            const noteDuration = timestamps.length === 1 ? 0 : timestamps[timestamps.length - 1] - timestamp;

            notesDetected.push({ at: timestamp, duration: noteDuration, value: note.step, octave: note.octave });
        }
    });
    if (notesDetected.length > 0) {
        // sort in ascending order based on notes timestamps
        notesDetected.sort((a: NotePlayed, b: NotePlayed) => a.at - b.at);
        console.log(notesDetected);

        // send notes to back
        apolloClient.mutate({
            mutation: AddNotesDocument,
            variables: {
                ID: perfID.value,
                inputNote: notesDetected,
            },
        });
    }

    if (stream.value !== null) {
        setTimeout(() => {
            getDuration(stream, apolloClient, perfID, startTimeStamp, acquisitionDelay);
        }, acquisitionDelay.value);
    } else {
        // no stream: end of performance
        apolloClient.mutate({
            mutation: FinishPerformanceDocument,
            variables: {
                performanceId: perfID.value,
            },
        });
    }
};
