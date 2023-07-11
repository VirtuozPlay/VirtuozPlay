/**
 * This code is the retranscription in typescript based on https://github.com/aerik/aerik.github.io
 */

import { Note, NotePlayed } from './notes';

let analyser: AnalyserNode;

let gainNode: GainNode;
let hertzBinSize: number;
let frequencyData: Uint8Array;
let buflen: number;

let notesPlayed: NotePlayed[] = []; // notes played during the track
let lastNotePlayed: NotePlayed = { octave: 0, step: '', timestamp: 0, duration: 0 };
let maxBytePlayed = 0; // max gain of note played
let startTimeStamp = 0;
// canvas
const maxValue = 256; //based on Uint8Array possible values

/**
 * Initialise settings of the microphone
 * @returns Audio stream
 */
export const initMicrophone = async (): Promise<MediaStream> => {
    const audioCtx: AudioContext = new window.AudioContext();

    analyser = audioCtx.createAnalyser();
    analyser.smoothingTimeConstant = 0.8; //default is 0.8, less is more responsive
    analyser.minDecibels = -45; //-100 is default and is more sensitive (more noise)
    analyser.fftSize = 8192 * 4; //need at least 8192 to detect differences in low notes

    const sampleRate: number = audioCtx.sampleRate;
    gainNode = audioCtx.createGain();
    gainNode.connect(audioCtx.destination);

    hertzBinSize = sampleRate / analyser.fftSize;
    frequencyData = new Uint8Array(analyser.frequencyBinCount);
    buflen = frequencyData.length;

    // ask for microphone permission
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true });

    const micSource: MediaStreamAudioSourceNode = audioCtx.createMediaStreamSource(stream);
    micSource.connect(gainNode);
    micSource.connect(analyser);

    startTimeStamp = Date.now();
    notesPlayed = [];

    return stream;
};

/**
 * Record played notes
 * @param notes         Array of notes associated with their frequencies
 * @param enableCanvas  Enable canvans
 * @param decibelMin    Minimum binary value of decibel to track (between 0 and 255)
 */
export const getTones = (notes: Note[], enableCanvas = false, decibelMin: number) => {
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
                notes[nPtr].power = total / count;
                // if a note is detected
                if (notes[nPtr].power > maxBytePlayed * 0.8 && notes[nPtr].power > decibelMin) {
                    maxBytePlayed = notes[nPtr].power;
                    const currentTimestamp = Date.now() - startTimeStamp;

                    // TODO check for note duration
                    // if(notesPlayed.length === 0) {
                    //     notesPlayed.push({
                    //         octave: notes[nPtr].octave,
                    //         step: notes[nPtr].step,
                    //         timestamp: currentTimestamp - startTimeStamp,
                    //         duration: 0,
                    //     });
                    // } else {
                    //     notesPlayed.forEach((notePlayed: NotePlayed) => {
                    //         if (
                    //             notes[nPtr].octave === notePlayed.octave &&
                    //             notes[nPtr].step === notePlayed.step &&
                    //             (currentTimestamp - startTimeStamp - 500) <= notePlayed.timestamp // if the note has been played in 500 ms
                    //         ) {
                    //             notesPlayed.push({
                    //                 octave: notes[nPtr].octave,
                    //                 step: notes[nPtr].step,
                    //                 timestamp: currentTimestamp - startTimeStamp,
                    //                 duration: currentTimestamp - notePlayed.timestamp,
                    //             });
                    //         }
                    //     });
                    // }
                    if (
                        // if note is different
                        notes[nPtr].octave !== lastNotePlayed.octave ||
                        notes[nPtr].step !== lastNotePlayed.step ||
                        // if the same note has been played after 500 ms
                        (notes[nPtr].octave === lastNotePlayed.octave &&
                            notes[nPtr].step === lastNotePlayed.step &&
                            currentTimestamp - 500 >= lastNotePlayed.timestamp)
                    ) {
                        notesPlayed.push({
                            octave: notes[nPtr].octave,
                            step: notes[nPtr].step,
                            timestamp: currentTimestamp,
                            duration: 0,
                        });

                        console.log(JSON.stringify(notes[nPtr]) + ' timestamp: ' + currentTimestamp);
                    }
                    // save the note that has been detected
                    lastNotePlayed = {
                        octave: notes[nPtr].octave,
                        step: notes[nPtr].step,
                        timestamp: currentTimestamp,
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
