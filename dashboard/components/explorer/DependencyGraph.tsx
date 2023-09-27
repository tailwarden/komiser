/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, memo } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import Cytoscape, { EventObject } from 'cytoscape';

import nodeHtmlLabel, {
  CytoscapeNodeHtmlParams
  // @ts-ignore
} from 'cytoscape-node-html-label';

// @ts-ignore
import COSEBilkent from 'cytoscape-cose-bilkent';

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
const DependencyGraph = ({ data }: DependencyGraphProps) => {
  const [initDone, setInitDone] = useState(false);

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
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9;" class="text-black-700 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center">${
                   templateData.label || '&nbsp;'
                 }</p>
              <p style="font-size: 10px; text-shadow: 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9,
                 0 0 5px #F4F9F9,0 0 5px #F4F9F9;" class="text-black-400 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center font-thin">${
                   templateData.service || '&nbsp;'
                 }</p></div>`;
          }
        }
      ]);
      // Add class to leave nodes so we can make them smaller
      cy.nodes().leaves().addClass('leaf');
      // same for root notes
      cy.nodes().roots().addClass('root');
      // Animate edges
      cy.edges().forEach(loopAnimation);

      // Hide labels when being zoomed out
      cy.on('zoom', event => {
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
    </div>
  );
};

export default memo(DependencyGraph);
