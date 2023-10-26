/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, memo, useEffect, useRef } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import Cytoscape, { EdgeSingular, EventObject } from 'cytoscape';
import popper from 'cytoscape-popper';

import nodeHtmlLabel, {
  CytoscapeNodeHtmlParams
  // @ts-ignore
} from 'cytoscape-node-html-label';

// @ts-ignore
import COSEBilkent from 'cytoscape-cose-bilkent';

import EmptyState from '@components/empty-state/EmptyState';

import Tooltip from '@components/tooltip/Tooltip';
import WarningIcon from '@components/icons/WarningIcon';
import PlusIcon from '@components/icons/PlusIcon';
import MinusIcon from '@components/icons/MinusIcon';
import SlashIcon from '@components/icons/SlashIcon';
import DragIcon from '@components/icons/DragIcon';
import { ReactFlowData } from './hooks/useDependencyGraph';
import {
  edgeAnimationConfig,
  edgeStyleConfig,
  graphLayoutConfig,
  leafStyleConfig,
  maxZoom,
  minZoom,
  nodeHTMLLabelConfig,
  nodeStyeConfig,
  // popperStyleConfig,
  zoomLevelBreakpoint
} from './config';

export type DependencyGraphProps = {
  data: ReactFlowData;
};

nodeHtmlLabel(Cytoscape.use(COSEBilkent));
Cytoscape.use(popper);
const DependencyGraph = ({ data }: DependencyGraphProps) => {
  const [initDone, setInitDone] = useState(false);
  const dataIsEmpty: boolean = data.nodes.length === 0;

  const [zoomLevel, setZoomLevel] = useState(minZoom);
  const [zoomVal, setZoomVal] = useState(minZoom); // debounced zoom state

  const percentageZoomChange = ((maxZoom - minZoom) / 100) * 5; // increase or decrease by 5%

  const [isNodeDraggingEnabled, setNodeDraggingEnabled] = useState(true);

  const cyRef = useRef<Cytoscape.Core | null>(null);

  // Type technically is Cytoscape.EdgeCollection but that throws an unexpected error
  const loopAnimation = (eles: any) => {
    const ani = eles.animation(edgeAnimationConfig[0], edgeAnimationConfig[1]);

    ani
      .reverse()
      .play()
      .promise('complete')
      .then(() => loopAnimation(eles));
  };

  const cyActionHandlers = (cy: Cytoscape.Core) => {
    // make sure we did not init already, otherwise this will be bound more than once
    if (!initDone) {
      // Add HTML labels for better flexibility
      // @ts-ignore
      cy.nodeHtmlLabel([
        {
          ...nodeHTMLLabelConfig,
          tpl(templateData: Cytoscape.NodeDataDefinition) {
            return `<div><p style="font-size: 10px; text-shadow: 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9;" class="text-black-700 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center" title="${
                   templateData.label
                 }">${templateData.label || '&nbsp;'}</p>
              <p style="font-size: 10px; text-shadow: 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9;" class="text-black-400 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center font-thin" title="${
                   templateData.label
                 }">${templateData.service || '&nbsp;'}</p></div>`;
          }
        }
      ]);
      // Add class to leave nodes so we can make them smaller
      cy.nodes().leaves().addClass('leaf');
      // same for root notes
      cy.nodes().roots().addClass('root');
      // Animate edges
      cy.edges().forEach(loopAnimation);

      // Add hover tooltip on edges
      cy.edges().bind('mouseover', event => {
        if (cy.zoom() >= zoomLevelBreakpoint) {
          // eslint-disable-next-line no-param-reassign
          event.target.popperRefObj = event.target.popper({
            content: () => {
              const content = document.createElement('div');
              content.classList.add('popper-div');
              content.innerHTML = event.target.data('label');
              content.style.pointerEvents = 'none';

              document.body.appendChild(content);
              return content;
            }
          });
        }
      });
      // Hide Edges tooltip on mouseout
      cy.edges().bind('mouseout', event => {
        if (cy.zoom() >= zoomLevelBreakpoint && event.target.popperRefObj) {
          event.target.popperRefObj.state.elements.popper.remove();
          event.target.popperRefObj.destroy();
        }
      });

      // Hide labels when being zoomed out
      cy.on('zoom', event => {
        const newZoomLevel = event.cy.zoom();
        // setZoomLevel(newZoomLevel);

        if (newZoomLevel <= zoomLevelBreakpoint) {
          interface ExtendedEdgeSingular extends EdgeSingular {
            popperRefObj?: any;
          }

          // Check if a tooltip is present and remove it
          cy.edges().forEach((edge: ExtendedEdgeSingular) => {
            if (edge.popperRefObj) {
              edge.popperRefObj.state.elements.popper.remove();
              edge.popperRefObj.destroy();
            }
          });
        }

        // update state with new zoom level
        setZoomLevel(newZoomLevel);

        const opacity = cy.zoom() <= zoomLevelBreakpoint ? 0 : 1;

        Array.from(
          document.querySelectorAll('.dependency-graph-nodeLabel'),
          e => {
            // @ts-ignore
            e.style.opacity = opacity;
            return e;
          }
        );
      });
      // Make sure to tell we inited successfully and prevent another init
      setInitDone(true);
    }
  };

  useEffect(() => {
    const handler = setTimeout(() => {
      setZoomVal(zoomLevel);
    }, 100); // 100ms debounce
    return () => {
      clearTimeout(handler);
    };
  }, [zoomLevel]);

  const toggleNodeDragging = () => {
    if (cyRef.current) {
      if (isNodeDraggingEnabled) {
        // to disable node dragging in Cytoscape
        cyRef.current.nodes().ungrabify();
      } else {
        // to enable node dragging in Cytoscape
        cyRef.current.nodes().grabify();
      }
      setNodeDraggingEnabled(!isNodeDraggingEnabled);
    }
  };

  const handleZoomChange = (zoomLevelNo: number) => {
    let newZoomLevel = zoomLevelNo;
    if (newZoomLevel < minZoom) newZoomLevel = minZoom;
    if (newZoomLevel > maxZoom) newZoomLevel = maxZoom;
    if (cyRef.current) {
      cyRef.current.zoom(newZoomLevel);
      setZoomLevel(newZoomLevel);
    }
  };

  return (
    <div className="relative h-full flex-1 bg-dependency-graph bg-[length:40px_40px]">
      {/* <CytoscapeComponent
        className="h-full w-full"
        elements={CytoscapeComponent.normalizeElements({
          nodes: data.nodes,
          edges: data.edges
        })}
        maxZoom={maxZoom}
        minZoom={minZoom}
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
        cy={(cy: Cytoscape.Core) => cyActionHandlers(cy)}
      /> */}
      {dataIsEmpty ? (
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
      <div className="absolute bottom-0 w-full">
        <div className="flex w-full flex-col items-center gap-2 sm:flex-row sm:justify-between">
          <div className="flex gap-2 overflow-visible bg-black-100 text-black-400">
            {data?.nodes?.length} Resources
            {!dataIsEmpty && (
              <div className="relative">
                <WarningIcon className="peer" height="16" width="16" />
                <Tooltip bottom="xs" align="left" width="lg">
                  Only AWS resources are currently supported on the explorer.
                </Tooltip>
              </div>
            )}
          </div>
          <div className="flex h-11 gap-4 overflow-visible bg-black-100 text-black-400">
            <button
              className="peer relative flex items-center rounded border bg-white p-3"
              onClick={toggleNodeDragging}
            >
              <DragIcon />
              {isNodeDraggingEnabled && (
                <SlashIcon className="absolute left-0 text-black-800" />
              )}
            </button>
            <Tooltip align="center" bottom="sm">
              {isNodeDraggingEnabled
                ? 'Disable node dragging'
                : 'Enable node dragging'}
            </Tooltip>

            <div className="flex">
              <button
                className="flex items-center rounded-l border bg-white p-3 shadow-sm"
                onClick={() => handleZoomChange(zoomVal - percentageZoomChange)}
              >
                <MinusIcon />
              </button>
              <div className="flex w-20 items-center justify-center border-b border-t bg-white px-5 py-3 shadow-sm">
                {Math.round(((zoomVal - minZoom) / (maxZoom - minZoom)) * 100)}%
              </div>
              <button
                className="flex items-center rounded-r border bg-white p-3 shadow-sm"
                onClick={() => handleZoomChange(zoomVal + percentageZoomChange)}
              >
                <PlusIcon />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default memo(DependencyGraph);
