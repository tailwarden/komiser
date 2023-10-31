/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, memo, useEffect } from 'react';
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
        if (cy.zoom() <= zoomLevelBreakpoint) {
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

        const opacity = cy.zoom() <= zoomLevelBreakpoint ? 0 : 1;

        Array.from(
          document.querySelectorAll('.dependency-graph-node-label'),
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

  return (
    <div className="relative h-full flex-1 bg-dependency-graph bg-[length:40px_40px]">
      <CytoscapeComponent
        className="h-full w-full"
        elements={CytoscapeComponent.normalizeElements({
          nodes: data.nodes,
          edges: data.edges
        })}
        // forEach={(data.nodes, node => {console.log(node)})}
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
      />
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
            cy={(cy: Cytoscape.Core) => cyActionHandlers(cy)}
          />
        </>
      )}
      <div className="absolute bottom-0 left-0 flex gap-2 overflow-visible bg-black-100 text-black-400">
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
