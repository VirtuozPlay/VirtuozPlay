/**
 * This code is the retranscription in typescript based on https://github.com/aerik/aerik.github.io
 */

import { notes } from './notes';

let analyser: AnalyserNode;

let gainNode: GainNode;
let hertzBinSize: number;
let frequencyData: Uint8Array;
let buflen: number;

const maxValue: number = 256; //based on Uint8Array possible values

const InitMicrophone = (enableCanvas: boolean = false) => {
    const audioCtx: AudioContext = new window.AudioContext();

    analyser = audioCtx.createAnalyser();
    analyser.smoothingTimeConstant = 0.2; //default is 0.8, less is more responsive
    analyser.minDecibels = -95; //-100 is default and is more sensitive (more noise)
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
const getTones = (enableCanvas: boolean = false) => {
    analyser.getByteFrequencyData(frequencyData);
    let count: number = 0;
    let total: number = 0;
    let sum: number = 0;
    const cutoff: number = 20; //redundant with decibels?
    let nPtr: number = 0; //notePointer
    for (let i = 0; i < buflen; i++) {
        const fdat: number = frequencyData[i];
        const freq: number = i * hertzBinSize; //freq in hertz for this sample
        const curNote = notes[nPtr];
        const nextNote = notes[nPtr + 1];
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
        const canvas: HTMLCanvasElement = <HTMLCanvasElement>document.getElementById('visualizer');
        const ctx: CanvasRenderingContext2D = canvas.getContext('2d')!;

        canvas.height = notes.length * 10;
        canvas.width = maxValue + 120;
        ctx.textAlign = 'left';
        //ctx.clearRect(0,0,canvas.width, canvas.height);
        for (let n = 0; n < notes.length; n++) {
            ctx.fillText(notes[n].note, 65, n * 10);
            ctx.save();
            ctx.textAlign = 'right';
            ctx.fillText(notes[n].frequency + ' Hz', 60, n * 10);
            ctx.restore();
        }
        ctx.save();
        //horzontal bars
        for (let n = 0; n < notes.length; n++) {
            const colString = 'hsl(' + (360 * n) / notes.length + ',100%,80%)';
            ctx.fillStyle = colString;
            ctx.fillRect(120, n * 10, notes[n].power!, -10);
        }
        ctx.restore();
    }
    setTimeout(() => {
        requestAnimationFrame(() => getTones(enableCanvas));
    }, 50);
};

export default InitMicrophone;
