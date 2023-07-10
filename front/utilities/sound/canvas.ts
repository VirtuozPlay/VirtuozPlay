import { Note } from './notes';

const maxValue = 256; //based on Uint8Array possible values

const showCanvas = (notes: Note[]) => {
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
        ctx.fillRect(120, n * 10, notes[n].power ?? 0, -10);
    }
    ctx.restore();
};

export default showCanvas;
