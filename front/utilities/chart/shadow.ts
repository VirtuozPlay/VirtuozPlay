import { AnimationEvent, Chart, ChartType, DoughnutControllerChartOptions, Plugin } from 'chart.js';
import { easingEffects, noop } from 'chart.js/helpers';

export interface ShadowOptions {
    enable?: boolean;
    /** Horizontal offset of the shadow from the chart X position. */
    offsetX?: number;
    /** Vertical offset of the shadow from the chart Y position. */
    offsetY?: number;
    animationDuration?: number;
    color?: string | CanvasGradient | CanvasPattern;
    /**
     * The size of the central hole in percents, 1.0 means covers the whole shape.
     * Defaults to the value of the `cutout` option of the chart.
     */
    cutout?: number;
    fillPercent?: number;
    /** Easing function to use for the rotating animation. */
    easing?: keyof typeof easingEffects;
}

type SupportedChartType = 'pie' | 'doughnut';

// Extend the valid configuration options for pie and doughnut charts with our shadow plugin options.
declare module 'chart.js' {
    export interface PluginOptionsByType<TType extends ChartType> {
        /** Options for the shadow plugin. */
        shadow?: false | (TType extends SupportedChartType ? ShadowOptions : never);
    }
}

/**
 *
 */
class _Shadow implements Plugin<SupportedChartType, ShadowOptions> {
    readonly id = 'shadow';

    private readonly _pieAnimationStarts: Map<string, number>;

    constructor() {
        this._pieAnimationStarts = new Map<string, number>();
    }

    /**
     * Default plugin options.
     */
    public defaults: Required<ShadowOptions> = {
        enable: false,
        offsetX: 16,
        offsetY: 16,
        animationDuration: 1000,
        color: 'black',
        cutout: 0.00001,
        fillPercent: 1.0,
        easing: 'easeOutQuart',
    };

    public beforeInit(chart: Chart, args: unknown, options: ShadowOptions) {
        if (!options || !options.enable) return;

        const animation = chart.options.animation;

        if (animation) {
            const oldOnProgress = (animation.onProgress || noop).bind(chart);
            const oldOnComplete = (animation.onComplete || noop).bind(chart);

            animation.onProgress = (event) => {
                oldOnProgress(event);
                this._onAnimationProgress(event);
            };
            animation.onComplete = (event) => {
                oldOnComplete(event);
                this._onAnimationComplete(event);
            };
        }

        if ('cutout' in chart.options) {
            const cutout = chart.options.cutout as DoughnutControllerChartOptions['cutout'];

            if (typeof cutout === 'string' && cutout.endsWith('%')) {
                options.cutout = parseFloat(cutout) / 100;
            }
        }
    }

    public beforeDraw(chart: Chart, args: unknown, options: ShadowOptions) {
        if (!options || !options.enable) return;
        // cast is safe, ChartJS merges defaults and user-provided options together
        const opts = options as Required<ShadowOptions>;
        const ctx: CanvasRenderingContext2D = chart.ctx;

        const width = chart.chartArea.right - chart.chartArea.left;
        const height = chart.chartArea.bottom - chart.chartArea.top;

        const centerX = width / 2 + opts.offsetX;
        const centerY = height / 2 + opts.offsetY;
        const radius = Math.min(width, height) / 2;
        const fill = opts.fillPercent * this._getEasedAnimationProgress(chart, opts);

        this._drawDoughnut(ctx, centerX, centerY, radius, fill, opts.cutout, opts.color);
    }

    private _onAnimationProgress(event: AnimationEvent) {
        if (event.initial && !this._pieAnimationStarts.has(event.chart.id)) {
            this._pieAnimationStarts.set(event.chart.id, Date.now());
        }
    }

    private _onAnimationComplete(event: AnimationEvent) {
        if (event.initial) {
            this._pieAnimationStarts.delete(event.chart.id);
        }
    }

    private _getEasedAnimationProgress(chart: Chart, options: Required<ShadowOptions>) {
        const pieAnimationStart = this._pieAnimationStarts.get(chart.id);

        if (pieAnimationStart != null) {
            const elapsed = Math.min(Date.now() - pieAnimationStart, options.animationDuration);
            return easingEffects[options.easing](elapsed / options.animationDuration);
        }
        return 1.0;
    }

    /**
     * Draws a doughnut shape using Canvas.
     */
    private _drawDoughnut(
        ctx: CanvasRenderingContext2D,
        cx: number,
        cy: number,
        radius: number,
        fillPercent: number,
        cutout: number,
        color: string | CanvasGradient | CanvasPattern
    ) {
        ctx.save();

        const startAngle = -Math.PI / 2; // start at top
        const fillAngle = startAngle + Math.PI * 2 * fillPercent;

        ctx.globalCompositeOperation = 'destination-over';
        ctx.fillStyle = color;
        ctx.beginPath();

        // line from inner to outer arc
        ctx.moveTo(cx, cy - radius * cutout);
        // outer arc
        ctx.arc(cx, cy, radius, startAngle, fillAngle, false);
        // line from outer to inner arc
        ctx.lineTo(Math.cos(fillAngle) * radius * cutout + cx, Math.sin(fillAngle) * radius * cutout + cy);
        // inner arc
        ctx.arc(cx, cy, radius * cutout, fillAngle, startAngle, true);

        ctx.fill();

        ctx.restore();
    }
}

/**
 * Chart.js plugin to add shadow to pie and doughnut charts.
 */
export const Shadow = new _Shadow();

export default Shadow;
