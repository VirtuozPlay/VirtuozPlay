/**
 * This code is the retranscription in typescript based on https://github.com/aerik/aerik.github.io
 */

import { notes, Note } from './notes';

let analyser: AnalyserNode;

let gainNode: GainNode;
let hertzBinSize: number;
let frequencyData: Uint8Array;
let buflen: number;

const maxValue = 256; //based on Uint8Array possible values

const InitMicrophone = (enableCanvas = false) => {
    const audioCtx: AudioContext = new window.AudioContext();

    analyser = audioCtx.createAnalyser();
    analyser.smoothingTimeConstant = 0.1; //default is 0.8, less is more responsive
    analyser.minDecibels = -60; //-100 is default and is more sensitive (more noise)
    analyser.fftSize = 8192 * 4; //need at least 8192 to detect differences in low notes

    const sampleRate: number = audioCtx.sampleRate;
    gainNode = audioCtx.createGain();
    gainNode.connect(audioCtx.destination);

    hertzBinSize = sampleRate / analyser.fftSize;
    console.log('bin size: ' + hertzBinSize);
    frequencyData = new Uint8Array(analyser.frequencyBinCount);
    buflen = frequencyData.length;

    // ask for mic permission
    navigator.mediaDevices
        .getUserMedia({ audio: true })
        .then((stream: MediaStream) => {
            const micSource: MediaStreamAudioSourceNode = audioCtx.createMediaStreamSource(stream);
            micSource.connect(gainNode);
            micSource.connect(analyser);

            getTones(enableCanvas);
        })
        .catch((err) => console.error(err)); // Disable error in console when blocking mic
};

//this basically lumps loud tones together and gets their avg frequency
const getTones = (enableCanvas = false) => {
    analyser.getByteFrequencyData(frequencyData);
    let count = 0;
    let total = 0;
    let sum = 0;
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
                sum += i; //bin numbers
                count++;
                total += fdat;
            }
        } else {
            if (count > 0) {
                const binNum: number = sum / count;
                //const bin = {};
                //round up
                let power: number = frequencyData[Math.ceil(binNum)];
                if (binNum > 0) {
                    //round down
                    power = (power + frequencyData[Math.floor(binNum)]) / 2;
                }
                //notes[nPtr].power = power;
                //seems like it rounds the values too much?
                notes[nPtr].power = total / count;
                sum = 0;
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
        showCanvas(notes);
    }
    setTimeout(() => {
        requestAnimationFrame(() => getTones(enableCanvas));
    }, 50);
};

export default InitMicrophone;
