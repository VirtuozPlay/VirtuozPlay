/**
 * This code is the retranscription in typescript based on https://github.com/aerik/aerik.github.io
 * No licence
 */

const notes: { note: string; frequency: number; power: number | undefined }[] = [
    //region Octave -1
    { note: 'do-1', frequency: 16.35, power: undefined },
    { note: 'do#-1', frequency: 17.32, power: undefined },
    { note: 're-1', frequency: 18.35, power: undefined },
    { note: 're#-1', frequency: 19.45, power: undefined },
    { note: 'mi-1', frequency: 20.6, power: undefined },
    { note: 'fa-1', frequency: 21.83, power: undefined },
    { note: 'fa#-1', frequency: 23.12, power: undefined },
    { note: 'sol-1', frequency: 24.5, power: undefined },
    { note: 'sol#-1', frequency: 25.96, power: undefined },
    { note: 'la-1', frequency: 27.5, power: undefined },
    { note: 'la#-1', frequency: 29.14, power: undefined },
    { note: 'si-1', frequency: 30.87, power: undefined },
    //endregion

    //region Octave 0
    { note: 'do0', frequency: 32.7, power: undefined },
    { note: 'do#0', frequency: 34.65, power: undefined },
    { note: 're0', frequency: 36.71, power: undefined },
    { note: 're#0', frequency: 38.89, power: undefined },
    { note: 'mi0', frequency: 41.2, power: undefined },
    { note: 'fa0', frequency: 43.65, power: undefined },
    { note: 'fa#0', frequency: 46.25, power: undefined },
    { note: 'sol0', frequency: 49, power: undefined },
    { note: 'sol#0', frequency: 51.91, power: undefined },
    { note: 'la0', frequency: 55, power: undefined },
    { note: 'la#0', frequency: 58.27, power: undefined },
    { note: 'si0', frequency: 61.74, power: undefined },
    //endregion

    //region Octave 1
    { note: 'do1', frequency: 65.41, power: undefined },
    { note: 'do#1', frequency: 69.3, power: undefined },
    { note: 're1', frequency: 73.42, power: undefined },
    { note: 're#1', frequency: 77.78, power: undefined },
    { note: 'mi1', frequency: 82.41, power: undefined },
    { note: 'fa1', frequency: 87.31, power: undefined },
    { note: 'fa#1', frequency: 92.5, power: undefined },
    { note: 'sol1', frequency: 98, power: undefined },
    { note: 'sol#1', frequency: 103.83, power: undefined },
    { note: 'la1', frequency: 110, power: undefined },
    { note: 'la#1', frequency: 116.54, power: undefined },
    { note: 'si1', frequency: 123.47, power: undefined },
    //endregion

    //region Octave 2
    { note: 'do2', frequency: 130.81, power: undefined },
    { note: 'do#2', frequency: 138.59, power: undefined },
    { note: 're2', frequency: 146.83, power: undefined },
    { note: 're#2', frequency: 155.56, power: undefined },
    { note: 'mi2', frequency: 164.81, power: undefined },
    { note: 'fa2', frequency: 174.61, power: undefined },
    { note: 'fa#2', frequency: 185, power: undefined },
    { note: 'sol2', frequency: 196, power: undefined },
    { note: 'sol#2', frequency: 207.65, power: undefined },
    { note: 'la2', frequency: 220, power: undefined },
    { note: 'la#2', frequency: 233.08, power: undefined },
    { note: 'si2', frequency: 246.94, power: undefined },
    //endregion

    //region Octave 3
    { note: 'do3', frequency: 261.63, power: undefined },
    { note: 'do#3', frequency: 277.18, power: undefined },
    { note: 're3', frequency: 293.66, power: undefined },
    { note: 're#3', frequency: 311.13, power: undefined },
    { note: 'mi3', frequency: 329.63, power: undefined },
    { note: 'fa3', frequency: 349.23, power: undefined },
    { note: 'fa#3', frequency: 369.99, power: undefined },
    { note: 'sol3', frequency: 392, power: undefined },
    { note: 'sol#3', frequency: 415.3, power: undefined },
    { note: 'la3', frequency: 440, power: undefined },
    { note: 'la#3', frequency: 466.16, power: undefined },
    { note: 'si3', frequency: 493.88, power: undefined },
    //endregion

    //region Octave 4
    { note: 'do4', frequency: 523.25, power: undefined },
    { note: 'do#4', frequency: 554.37, power: undefined },
    { note: 're4', frequency: 587.33, power: undefined },
    { note: 're#4', frequency: 622.25, power: undefined },
    { note: 'mi4', frequency: 659.25, power: undefined },
    { note: 'fa4', frequency: 698.46, power: undefined },
    { note: 'fa#4', frequency: 739.99, power: undefined },
    { note: 'sol4', frequency: 783.99, power: undefined },
    { note: 'sol#4', frequency: 830.61, power: undefined },
    { note: 'la4', frequency: 880, power: undefined },
    { note: 'la#4', frequency: 932.33, power: undefined },
    { note: 'si4', frequency: 987.77, power: undefined },
    //endregion

    //region Octave 5
    { note: 'do5', frequency: 1046.5, power: undefined },
    { note: 'do#5', frequency: 1108.73, power: undefined },
    { note: 're5', frequency: 1174.66, power: undefined },
    { note: 're#5', frequency: 1244.51, power: undefined },
    { note: 'mi5', frequency: 1318.51, power: undefined },
    { note: 'fa5', frequency: 1396.91, power: undefined },
    { note: 'fa#5', frequency: 1479.98, power: undefined },
    { note: 'sol5', frequency: 1567.98, power: undefined },
    { note: 'sol#5', frequency: 1661.22, power: undefined },
    { note: 'la5', frequency: 1760, power: undefined },
    { note: 'la#5', frequency: 1864.66, power: undefined },
    { note: 'si5', frequency: 1975.53, power: undefined },
    //endregion

    //region Octave 6
    { note: 'do6', frequency: 2093, power: undefined },
    { note: 'do#6', frequency: 2217.46, power: undefined },
    { note: 're6', frequency: 2349.32, power: undefined },
    { note: 're#6', frequency: 2489.02, power: undefined },
    { note: 'mi6', frequency: 2637.02, power: undefined },
    { note: 'fa6', frequency: 2793.83, power: undefined },
    { note: 'fa#6', frequency: 2959.96, power: undefined },
    { note: 'sol6', frequency: 3135.96, power: undefined },
    { note: 'sol#6', frequency: 3322.44, power: undefined },
    { note: 'la6', frequency: 3520, power: undefined },
    { note: 'la#6', frequency: 3729.31, power: undefined },
    { note: 'si6', frequency: 3951.07, power: undefined },
    //endregion

    //region Octave 7
    { note: 'do7', frequency: 4186.01, power: undefined },
    { note: 'do#7', frequency: 4434.92, power: undefined },
    { note: 're7', frequency: 4698.63, power: undefined },
    { note: 're#7', frequency: 4978.03, power: undefined },
    { note: 'mi7', frequency: 5274.04, power: undefined },
    { note: 'fa7', frequency: 5587.65, power: undefined },
    { note: 'fa#7', frequency: 5919.91, power: undefined },
    { note: 'sol7', frequency: 6271.93, power: undefined },
    { note: 'sol#7', frequency: 6644.88, power: undefined },
    { note: 'la7', frequency: 7040, power: undefined },
    { note: 'la#7', frequency: 7458.62, power: undefined },
    { note: 'si7', frequency: 7902.13, power: undefined },
    //endregion
];

//const maxValue: number = 256; //based on Uint8Array possible values
let analyser: AnalyserNode;

let gainNode: GainNode;
let hertzBinSize: number;
let frequencyData: Uint8Array;
let buflen: number;

function initMicrophone() {
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

    navigator.mediaDevices.getUserMedia({ audio: true }).then((stream: MediaStream) => {
        const micSource: MediaStreamAudioSourceNode = audioCtx.createMediaStreamSource(stream);
        micSource.connect(gainNode);
        micSource.connect(analyser);

        getTones();
    });
}

//this basically lumps loud tones together and gets their avg frequency
function getTones() {
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
                const binNum = sum / count;
                //const bin = {};
                //round up
                let power = frequencyData[Math.ceil(binNum)];
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
    //drawNotes();
    setTimeout(function () {
        requestAnimationFrame(getTones);
    }, 50);
}
