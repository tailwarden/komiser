import { Core, EdgeSingular, NodeSingular } from 'cytoscape';
import { IPoint } from 'cytoscape-layers';
import { edgeAnimationConfig, edgeStyleConfig } from './config';

// maybe this can be needed in the future
export function dashAnimation(
  edge: EdgeSingular,
  duration: number,
  offset: number
) {
  return edge
    .animation({
      style: {
        'line-dash-offset': offset,
        'line-dash-pattern': edgeAnimationConfig.style['line-dash-pattern']
      },
      duration,
      position: edge.sourceEndpoint(),
      renderedPosition: edge.sourceEndpoint(),
      easing: edgeAnimationConfig.easing as 'linear'
    })
    .play()
    .promise('complete');
}

type Options = {
  direction: 'alternate' | 'forward' | 'backward';
  mode: 'speed' | 'duration';
  modeValue: number;
  randomOffset: boolean;
};

function getOrSet<T>(
  elem: NodeSingular | EdgeSingular,
  key: string,
  value: () => T
): T {
  const v = elem.scratch(key);
  if (v != null) {
    return v;
  }
  const vSet = value();
  elem.scratch(key, vSet);
  return vSet;
}

function dist(start: IPoint, end: IPoint) {
  return Math.sqrt((start.x - end.x) ** 2 + (start.y - end.y) ** 2);
}

export const animateEdges = async (options: Options, cy: Core) => {
  let cyLayers;

  function computeFactor(
    elapsed: number,
    offset: number,
    start: IPoint,
    end: IPoint
  ) {
    const distance = dist(start, end);
    let duration = options.modeValue;

    if (options.mode !== 'duration' && options.modeValue !== 0) {
      duration = distance / options.modeValue;
    }

    if (
      !Number.isFinite(duration) ||
      Number.isNaN(duration) ||
      duration === 0
    ) {
      return 0;
    }

    let f = elapsed / duration;

    if (options.direction === 'alternate') {
      f = f / 2 + offset;
      const v = 2 * (f - Math.floor(f) - 0.5);
      return Math.abs(v);
    }

    f += offset;
    const v = f - Math.floor(f);
    return options.direction === 'forward' ? v : 1 - v;
  }

  if (typeof window !== 'undefined') {
    const cytoLayersModule = await import('cytoscape-layers');
    cyLayers = cytoLayersModule.layers(cy);
    const animationLayer = cyLayers.nodeLayer.insertBefore('canvas');

    let start: number | null = null;
    let elapsed = 0;
    const update = (time: number) => {
      if (start == null) {
        start = time;
      }
      elapsed = time - start;
      animationLayer.update();
      requestAnimationFrame(update);
    };
    cyLayers.renderPerEdge(
      animationLayer,
      (
        ctx: CanvasRenderingContext2D,
        edge: EdgeSingular,
        path: Path2D,
        startPoint: IPoint,
        end: IPoint
      ) => {
        const offset = options.randomOffset
          ? getOrSet(edge, '_animOffset', () => Math.random())
          : 0;
        const g = ctx.createLinearGradient(
          startPoint.x,
          startPoint.y,
          end.x,
          end.y
        );

        const v = computeFactor(elapsed, offset, startPoint, end);

        const colorStop1 = edgeStyleConfig['line-gradient-stop-colors']
          ? edgeStyleConfig['line-gradient-stop-colors'][0]
          : '#008484';
        const colorStop2 = edgeStyleConfig['line-gradient-stop-colors']
          ? edgeStyleConfig['line-gradient-stop-colors'][1]
          : '#33CCCC';

        if (typeof colorStop1 === 'string' && typeof colorStop2 === 'string') {
          g.addColorStop(Math.max(v - 0.1, 0), colorStop1);
          g.addColorStop(v, 'white');
          g.addColorStop(Math.min(v + 0.1, 1), colorStop2);
        }
        ctx.strokeStyle = g;
        ctx.lineWidth = 3;
        ctx.stroke(path);
      },
      {
        checkBounds: true,
        checkBoundsPointCount: 5
      }
    );
    requestAnimationFrame(update);
  }
};
