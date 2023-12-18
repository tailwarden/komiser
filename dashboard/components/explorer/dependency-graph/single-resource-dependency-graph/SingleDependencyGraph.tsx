import React, { memo } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import Cytoscape from 'cytoscape';

import EmptyState from '@components/empty-state/EmptyState';

import Tooltip from '@components/tooltip/Tooltip';
import { DragIcon } from '@components/icons';
import NumberInput from '@components/number-input/NumberInput';
import { DependencyGraphProps } from '../hooks/useDependencyGraph';
import {
  edgeStyleConfig,
  graphLayoutConfig,
  leafStyleConfig,
  maxZoom,
  minZoom,
  nodeStyeConfig
} from '../config';
import { useDependencyGraphActions } from '../hooks/useDependencyGraphActions';

const SingleDependencyGraph = ({ data }: DependencyGraphProps) => {
  const dataNodesLength: number = data.nodes.length;
  const {
    cyRef,
    cyActionHandlers,
    toggleNodeDragging,
    isNodeDraggingEnabled,
    translateXClass,
    zoomVal,
    handleZoomChange
  } = useDependencyGraphActions({ isSingleDependencyGraph: true });

  return (
    <div className="relative h-full flex-1 bg-dependency-graph bg-[length:40px_40px]">
      {dataNodesLength === 0 ? (
        <>
          <div className="translate-y-[201px]">
            <EmptyState
              title="No results for this filter"
              message="It seems like you have no cloud resources matching the filters you added"
              mascotPose="tablet"
            />
          </div>
        </>
      ) : (
        <>
          <CytoscapeComponent
            className="h-full w-full"
            elements={CytoscapeComponent.normalizeElements({
              nodes: data.nodes,
              edges: data.edges
            })}
            maxZoom={maxZoom}
            minZoom={minZoom}
            zoom={(minZoom + maxZoom) / 2}
            layout={graphLayoutConfig}
            stylesheet={[
              {
                selector: 'node',
                style: nodeStyeConfig
              },
              {
                selector: 'edge',
                style: edgeStyleConfig
              },
              {
                selector: '.leaf',
                style: leafStyleConfig
              }
            ]}
            cy={(cy: Cytoscape.Core) => {
              cyActionHandlers(cy);
              cyRef.current = cy;
            }}
          />
        </>
      )}
      <div className="absolute bottom-5 right-1 w-full">
        <div className="flex w-full flex-col items-center gap-2 sm:flex-row sm:justify-between">
          <div className="ml-1 flex overflow-visible bg-black-100 text-black-400">
            {dataNodesLength}{' '}
            {`related resource${dataNodesLength > 1 ? 's' : ''}`}
          </div>
          <div className="flex max-h-11 gap-4">
            <button
              className={`peer relative flex items-center rounded border-[1.2px] border-black-200 bg-white p-2.5 ${
                isNodeDraggingEnabled && 'border-primary'
              }`}
              onClick={toggleNodeDragging}
            >
              <DragIcon className="h-6 w-6" />
            </button>
            <Tooltip align="center" bottom="sm">
              {isNodeDraggingEnabled
                ? 'Disable node dragging'
                : 'Enable node dragging'}
            </Tooltip>

            <div className="relative w-40">
              <NumberInput
                name="zoom"
                value={zoomVal}
                action={zoomData => handleZoomChange(Number(zoomData.zoom))}
                handleValueChange={handleZoomChange} // increment or decrement input value
                step={5} // percentage change in zoom
                maxLength={3}
              />
              <span
                className={`absolute left-1/2 top-1/2 ${translateXClass} -translate-y-1/2 text-sm text-neutral-900`}
              >
                %
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default memo(SingleDependencyGraph);
