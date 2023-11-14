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
import DragIcon from '@components/icons/DragIcon';
import NumberInput from '@components/number-input/NumberInput';
import useInventory from '@components/inventory/hooks/useInventory/useInventory';
import settingsService from '@services/settingsService';
import InventorySidePanel from '@components/inventory/components/InventorySidePanel';
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
  const [zoomVal, setZoomVal] = useState(0); // debounced zoom state to display percentage

  const [isNodeDraggingEnabled, setNodeDraggingEnabled] = useState(true);

  const cyRef = useRef<Cytoscape.Core | null>(null);
  const {
    openModal,
    isOpen,
    closeModal,
    data: inventoryItem,
    page,
    goTo,
    tags,
    handleChange,
    addNewTag,
    removeTag,
    updateTags,
    loading,
    deleteLoading,
    bulkItems,
    updateBulkTags
  } = useInventory();

  // opens modal to display details of clicked node
  const handleNodeClick = async (event: EventObject) => {
    const nodeData = event.target.data();
    settingsService.getResourceById(`?resourceId=${nodeData.id}`).then(res => {
      if (res !== Error) {
        openModal(res);
      }
    });
  };

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
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9;" class="text-gray-900 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center" title="${
                   templateData.label
                 }">${templateData.label || '&nbsp;'}</p>
              <p style="font-size: 10px; text-shadow: 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9;" class="text-gray-700 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center font-thin" title="${
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
      // Add a click event listener to the Cytoscape graph
      cy.on('tap', 'node', handleNodeClick);

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
    const zoomPercentage = Math.round(
      ((zoomLevel - minZoom) / (maxZoom - minZoom)) * 100
    );
    const handler = setTimeout(() => {
      setZoomVal(zoomPercentage);
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

  const handleZoomChange = (zoomPercentage: number) => {
    let newZoomLevel = minZoom + zoomPercentage * ((maxZoom - minZoom) / 100);
    if (newZoomLevel < minZoom) newZoomLevel = minZoom;
    if (newZoomLevel > maxZoom) newZoomLevel = maxZoom;
    if (cyRef.current) {
      cyRef.current.zoom(newZoomLevel);
      setZoomLevel(newZoomLevel);
    }
  };

  let translateXClass;

  if (zoomVal < 10) {
    translateXClass = 'translate-x-1';
  } else if (zoomVal >= 10 && zoomVal < 100) {
    translateXClass = 'translate-x-2';
  } else {
    translateXClass = 'translate-x-3';
  }

  return (
    <div className="relative h-full flex-1 bg-dependency-graph bg-[length:40px_40px]">
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
          <div className="flex items-center gap-2 bg-black-100 text-black-400">
            <span>{data?.nodes?.length} Resources</span>
            {!dataIsEmpty && (
              <div className="relative mt-[2px]">
                <WarningIcon className="peer" height="16" width="16" />
                <Tooltip bottom="xs" align="left" width="lg">
                  Only AWS and CIVO resources are currently supported on the
                  explorer.
                </Tooltip>
              </div>
            )}
          </div>
          <div className="flex max-h-11 gap-4">
            <button
              className={`peer relative flex items-center rounded border-[1.2px] border-gray-300 bg-white p-2.5 ${
                isNodeDraggingEnabled && 'border-darkcyan-500'
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
                className={`absolute left-1/2 top-1/2 ${translateXClass} text-neutral-900 -translate-y-1/2 text-sm`}
              >
                %
              </span>
            </div>
          </div>
        </div>
      </div>
      {/* Modal */}
      <InventorySidePanel
        isOpen={isOpen}
        closeModal={closeModal}
        data={inventoryItem}
        goTo={goTo}
        page={page}
        updateTags={updateTags}
        tags={tags}
        tabs={['resource details', 'tags']}
        handleChange={handleChange}
        removeTag={removeTag}
        addNewTag={addNewTag}
        loading={loading}
        deleteLoading={deleteLoading}
        bulkItems={bulkItems}
        updateBulkTags={updateBulkTags}
      />
    </div>
  );
};

export default memo(DependencyGraph);
