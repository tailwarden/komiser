// useDependencyGraphActions.js
import { useEffect, useRef, useState } from 'react';
import Cytoscape, { EdgeSingular, EventObject } from 'cytoscape';
import popper from 'cytoscape-popper';
import nodeHtmlLabel, {
  CytoscapeNodeHtmlParams
  // @ts-ignore
} from 'cytoscape-node-html-label';
// @ts-ignore
import COSEBilkent from 'cytoscape-cose-bilkent';
import settingsService from '@services/settingsService';
import { InventoryItem } from '@components/inventory/hooks/useInventory/types/useInventoryTypes';
import {
  edgeAnimationConfig,
  maxZoom,
  minZoom,
  nodeHTMLLabelConfig,
  zoomLevelBreakpoint
} from '../config';

type UseDependencyGraphActionsT = {
  isSingleDependencyGraph?: boolean;
  openModal?: (inventoryItem: InventoryItem) => void;
};

// Register the extensions only once
nodeHtmlLabel(Cytoscape.use(COSEBilkent));
Cytoscape.use(popper);

export const useDependencyGraphActions = ({
  isSingleDependencyGraph = false,
  openModal
}: UseDependencyGraphActionsT) => {
  const [initDone, setInitDone] = useState(false);

  const [isNodeDraggingEnabled, setNodeDraggingEnabled] = useState(true);

  const [zoomLevel, setZoomLevel] = useState(minZoom);
  const [zoomVal, setZoomVal] = useState(0); // debounced zoom state to display percentage
  const cyRef = useRef<Cytoscape.Core | null>(null);
  const resourceId = JSON.parse(localStorage.getItem('resourceId') || '');

  const [zoomToResourceId, setZoomToResourceId] = useState(false);

  // opens modal to display details of clicked node
  const handleNodeClick = async (event: EventObject) => {
    const nodeData = event.target.data();
    settingsService.getResourceById(`?resourceId=${nodeData.id}`).then(res => {
      if (res !== Error) {
        if (openModal) openModal(res);
      }
    });
  };

  const loopAnimation = (eles: any) => {
    const ani = eles.animation(edgeAnimationConfig[0], edgeAnimationConfig[1]);
    ani
      .reverse()
      .play()
      .promise('complete')
      .then(() => loopAnimation(eles));
  };

  const cyActionHandlers = (cy: Cytoscape.Core) => {
    if (!initDone) {
      // Add HTML labels for better flexibility
      // @ts-ignore
      cy.nodeHtmlLabel([
        {
          ...nodeHTMLLabelConfig,
          tpl: (
            templateData: Cytoscape.NodeDataDefinition
          ) => `<div><p style="font-size: 10px; text-shadow: 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
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
                 }">${templateData.service || '&nbsp;'}</p></div>`
        }
      ]);

      // Add class to leave nodes so we can make them smaller
      cy.nodes().leaves().addClass('leaf');
      // same for root nodes
      cy.nodes().roots().addClass('root');
      // Animate edges
      cy.edges().forEach(loopAnimation);

      // Add a click event listener to the Cytoscape graph
      if (!isSingleDependencyGraph) {
        cy.on('tap', 'node', handleNodeClick);
      }

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

      if (resourceId && cyRef.current) {
        const targetNode = cyRef.current.getElementById(resourceId);

        if (targetNode.length > 0) {
          cyRef.current.fit(targetNode);
          setZoomToResourceId(true);
        }
      }
    }
  };
  useEffect(() => {
    if (cyRef.current && zoomToResourceId) {
      const targetNode = cyRef.current.getElementById(resourceId);

      if (targetNode.length > 0) {
        cyRef.current.fit(targetNode);
        setZoomToResourceId(false); // Reset the state after zooming
      }
    }
  }, [zoomToResourceId, cyRef.current]);

  const handleZoomChange = (zoomPercentage: number) => {
    let newZoomLevel = minZoom + zoomPercentage * ((maxZoom - minZoom) / 100);
    if (newZoomLevel < minZoom) newZoomLevel = minZoom;
    if (newZoomLevel > maxZoom) newZoomLevel = maxZoom;
    if (cyRef.current) {
      cyRef.current.zoom(newZoomLevel);
      setZoomLevel(newZoomLevel);
    }
  };

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

  let translateXClass;

  if (zoomVal < 10) {
    translateXClass = 'translate-x-1';
  } else if (zoomVal >= 10 && zoomVal < 100) {
    translateXClass = 'translate-x-2';
  } else {
    translateXClass = 'translate-x-3';
  }
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

  return {
    cyRef,
    cyActionHandlers,
    toggleNodeDragging,
    isNodeDraggingEnabled,
    handleZoomChange,
    zoomVal,
    translateXClass
  };
};
