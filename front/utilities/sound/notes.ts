export interface Note {
    octave: number;
    step: string;
    frequency: number;
    power?: number;
}

export const notes: Note[] = [
    //region Octave 0
    { octave: 0, step: 'C', frequency: 16.35, power: undefined },
    { octave: 0, step: 'C#/Db', frequency: 17.32, power: undefined },
    { octave: 0, step: 'D', frequency: 18.35, power: undefined },
    { octave: 0, step: 'D#/Eb', frequency: 19.45, power: undefined },
    { octave: 0, step: 'E', frequency: 20.6, power: undefined },
    { octave: 0, step: 'F', frequency: 21.83, power: undefined },
    { octave: 0, step: 'F#/Gb', frequency: 23.12, power: undefined },
    { octave: 0, step: 'G', frequency: 24.5, power: undefined },
    { octave: 0, step: 'G#/Ab', frequency: 25.96, power: undefined },
    { octave: 0, step: 'A', frequency: 27.5, power: undefined },
    { octave: 0, step: 'A#/Bb', frequency: 29.14, power: undefined },
    { octave: 0, step: 'B', frequency: 30.87, power: undefined },
    //endregion

    //region Octave 1
    { octave: 1, step: 'C', frequency: 32.7, power: undefined },
    { octave: 1, step: 'C#/Db', frequency: 34.65, power: undefined },
    { octave: 1, step: 'D', frequency: 36.71, power: undefined },
    { octave: 1, step: 'D#/Eb', frequency: 38.89, power: undefined },
    { octave: 1, step: 'E', frequency: 41.2, power: undefined },
    { octave: 1, step: 'F', frequency: 43.65, power: undefined },
    { octave: 1, step: 'F#/Gb', frequency: 46.25, power: undefined },
    { octave: 1, step: 'G', frequency: 49, power: undefined },
    { octave: 1, step: 'G#/Ab', frequency: 51.91, power: undefined },
    { octave: 1, step: 'A', frequency: 55, power: undefined },
    { octave: 1, step: 'A#/Bb', frequency: 58.27, power: undefined },
    { octave: 1, step: 'B', frequency: 61.74, power: undefined },
    //endregion

    //region Octave 2
    { octave: 2, step: 'C',  frequency: 65.41, power: undefined },
    { octave: 2, step: 'C#/Db', frequency: 69.3, power: undefined },
    { octave: 2, step: 'D',  frequency: 73.42, power: undefined },
    { octave: 2, step: 'D#/Eb', frequency: 77.78, power: undefined },
    { octave: 2, step: 'E',  frequency: 82.41, power: undefined },
    { octave: 2, step: 'F',  frequency: 87.31, power: undefined },
    { octave: 2, step: 'F#/Gb', frequency: 92.5, power: undefined },
    { octave: 2, step: 'G',  frequency: 98, power: undefined },
    { octave: 2, step: 'G#/Ab', frequency: 103.83, power: undefined },
    { octave: 2, step: 'A',  frequency: 110, power: undefined },
    { octave: 2, step: 'A#/Bb', frequency: 116.54, power: undefined },
    { octave: 2, step: 'B',  frequency: 123.47, power: undefined },
    //endregion

    //region Octave 3
    { octave: 3, step: 'C', frequency: 130.81, power: undefined },
    { octave: 3, step: 'C#/Db', frequency: 138.59, power: undefined },
    { octave: 3, step: 'D', frequency: 146.83, power: undefined },
    { octave: 3, step: 'D#/Eb', frequency: 155.56, power: undefined },
    { octave: 3, step: 'E', frequency: 164.81, power: undefined },
    { octave: 3, step: 'F', frequency: 174.61, power: undefined },
    { octave: 3, step: 'F#/Gb', frequency: 185, power: undefined },
    { octave: 3, step: 'G', frequency: 196, power: undefined },
    { octave: 3, step: 'G#/Ab', frequency: 207.65, power: undefined },
    { octave: 3, step: 'A', frequency: 220, power: undefined },
    { octave: 3, step: 'A#/Bb', frequency: 233.08, power: undefined },
    { octave: 3, step: 'B', frequency: 246.94, power: undefined },
    //endregion

    //region Octave 4
    { octave: 4, step: 'C', frequency: 261.63, power: undefined },
    { octave: 4, step: 'C#/Db', frequency: 277.18, power: undefined },
    { octave: 4, step: 'D', frequency: 293.66, power: undefined },
    { octave: 4, step: 'D#/Eb', frequency: 311.13, power: undefined },
    { octave: 4, step: 'E', frequency: 329.63, power: undefined },
    { octave: 4, step: 'F', frequency: 349.23, power: undefined },
    { octave: 4, step: 'F#/Gb', frequency: 369.99, power: undefined },
    { octave: 4, step: 'G', frequency: 392, power: undefined },
    { octave: 4, step: 'G#/Ab', frequency: 415.3, power: undefined },
    { octave: 4, step: 'A', frequency: 440, power: undefined },
    { octave: 4, step: 'A#/Bb', frequency: 466.16, power: undefined },
    { octave: 4, step: 'B', frequency: 493.88, power: undefined },
    //endregion

    //region Octave 5
    { octave: 5, step: 'C', frequency: 523.25, power: undefined },
    { octave: 5, step: 'C#/Db', frequency: 554.37, power: undefined },
    { octave: 5, step: 'D', frequency: 587.33, power: undefined },
    { octave: 5, step: 'D#/Eb', frequency: 622.25, power: undefined },
    { octave: 5, step: 'E', frequency: 659.25, power: undefined },
    { octave: 5, step: 'F', frequency: 698.46, power: undefined },
    { octave: 5, step: 'F#/Gb', frequency: 739.99, power: undefined },
    { octave: 5, step: 'G', frequency: 783.99, power: undefined },
    { octave: 5, step: 'G#/Ab', frequency: 830.61, power: undefined },
    { octave: 5, step: 'A', frequency: 880, power: undefined },
    { octave: 5, step: 'A#/Bb',  frequency: 932.33, power: undefined },
    { octave: 5, step: 'B', frequency: 987.77, power: undefined },
    //endregion

    //region Octave 6
    { octave: 6, step: 'C', frequency: 1046.5, power: undefined },
    { octave: 6, step: 'C#/Db', frequency: 1108.73, power: undefined },
    { octave: 6, step: 'D', frequency: 1174.66, power: undefined },
    { octave: 6, step: 'D#/Eb', frequency: 1244.51, power: undefined },
    { octave: 6, step: 'E', frequency: 1318.51, power: undefined },
    { octave: 6, step: 'F', frequency: 1396.91, power: undefined },
    { octave: 6, step: 'F#/Gb', frequency: 1479.98, power: undefined },
    { octave: 6, step: 'G', frequency: 1567.98, power: undefined },
    { octave: 6, step: 'G#/', frequency: 1661.22, power: undefined },
    { octave: 6, step: 'A', frequency: 1760, power: undefined },
    { octave: 6, step: 'A#/Bb', frequency: 1864.66, power: undefined },
    { octave: 6, step: 'B', frequency: 1975.53, power: undefined },
    //endregion

    //region Octave 7
    { octave: 7, step: 'C', frequency: 2093, power: undefined },
    { octave: 7, step: 'C#/Db', frequency: 2217.46, power: undefined },
    { octave: 7, step: 'D', frequency: 2349.32, power: undefined },
    { octave: 7, step: 'D#/Eb', frequency: 2489.02, power: undefined },
    { octave: 7, step: 'E', frequency: 2637.02, power: undefined },
    { octave: 7, step: 'F', frequency: 2793.83, power: undefined },
    { octave: 7, step: 'F#/Gb', frequency: 2959.96, power: undefined },
    { octave: 7, step: 'G', frequency: 3135.96, power: undefined },
    { octave: 7, step: 'G#/Ab', frequency: 3322.44, power: undefined },
    { octave: 7, step: 'A', frequency: 3520, power: undefined },
    { octave: 7, step: 'A#/Bb', frequency: 3729.31, power: undefined },
    { octave: 7, step: 'B', frequency: 3951.07, power: undefined },
    //endregion

    //region Octave 8
    { octave: 8, step: 'C', frequency: 4186.01, power: undefined },
    { octave: 8, step: 'C#/Db', frequency: 4434.92, power: undefined },
    { octave: 8, step: 'D', frequency: 4698.63, power: undefined },
    { octave: 8, step: 'D#/Eb', frequency: 4978.03, power: undefined },
    { octave: 8, step: 'E', frequency: 5274.04, power: undefined },
    { octave: 8, step: 'F', frequency: 5587.65, power: undefined },
    { octave: 8, step: 'F#/Gb', frequency: 5919.91, power: undefined },
    { octave: 8, step: 'G', frequency: 6271.93, power: undefined },
    { octave: 8, step: 'G#/Ab', frequency: 6644.88, power: undefined },
    { octave: 8, step: 'A', frequency: 7040, power: undefined },
    { octave: 8, step: 'A#/Bb', frequency: 7458.62, power: undefined },
    { octave: 8, step: 'B', frequency: 7902.13, power: undefined },
    //endregion
];
