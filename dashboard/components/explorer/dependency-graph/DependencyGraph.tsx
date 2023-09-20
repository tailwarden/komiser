/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, memo } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import Cytoscape, { EventObject } from 'cytoscape';

import nodeHtmlLabel, {
  CytoscapeNodeHtmlParams
} from 'cytoscape-node-html-label';

// @ts-ignore
import COSEBilkent from 'cytoscape-cose-bilkent';

import { ReactFlowData } from './hooks/useDependencyGraph';

export type DependencyGraphProps = {
  data: ReactFlowData;
};

const layout = { name: 'cose-bilkent', animate: true };

nodeHtmlLabel(Cytoscape.use(COSEBilkent));
const DependencyGraph = ({ data }: DependencyGraphProps) => {
  const [initDone, setInitDone] = useState(false);

  const cyActionHandlers = (cy: Cytoscape.Core) => {
    if (!initDone) {
      cy.on('click', 'node', (evt: EventObject) => {
        console.info(`Clicked the node with ID ${evt.target.id()}`);
      });
      // @ts-ignore
      cy.nodeHtmlLabel([
        {
          query: 'node', // cytoscape query selector
          halign: 'center', // title vertical position. Can be 'left',''center, 'right'
          valign: 'bottom', // title vertical position. Can be 'top',''center, 'bottom'
          halignBox: 'center', // title vertical position. Can be 'left',''center, 'right'
          valignBox: 'bottom', // title relative box vertical position. Can be 'top',''center, 'bottom'
          cssClass: '', // any classes will be as attribute of <div> container for every title
          tpl(templateData: Cytoscape.NodeDataDefinition) {
            return `<div>
              <p style="font-size: 10px;" class="text-black-700 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center">${
                templateData.label || '&nbsp;'
              }</p>
              <p style="font-size: 10px;" class="text-black-400 text-ellipsis max-w-[100px] overflow-hidden whitespace-nowrap text-center font-thin">${
                templateData.service || '&nbsp;'
              }</p>
            </div>`; // your html template here
          }
        }
      ]);
      cy.nodes().leaves().addClass('leaf');
      setInitDone(true);
    }
  };

  return (
    <div className="relative h-full flex-1 bg-dependency-graph bg-[length:40px_40px]">
      <div className="h-[70vh]">
        <CytoscapeComponent
          className="h-full w-full"
          elements={CytoscapeComponent.normalizeElements({
            nodes: data.nodes,
            edges: data.edges
          })}
          maxZoom={4}
          minZoom={0.25}
          layout={layout}
          stylesheet={[
            {
              selector: 'node',
              style: {
                width: 45,
                height: 45,
                shape: 'ellipse',
                // content: 'data(label)',
                'background-color': 'white',
                'background-image': node =>
                  node.data('provider') === 'AWS'
                    ? '/assets/img/dependency-graph/aws-node.svg'
                    : '',
                'background-height': 20,
                'border-color': '#EDEBEE',
                'border-width': 1,
                'border-style': 'solid'
              }
            },
            {
              selector: 'edge',
              style: {
                width: 1,
                'line-color': '#008484',
                'line-style': edge =>
                  edge.data('relation') === 'USES' ? 'solid' : 'dashed',
                'curve-style': 'unbundled-bezier',
                'control-point-distances': [8, -8],
                'control-point-weights': [0.15, 0.85]
              }
            },
            {
              selector: '.leaf',
              style: {
                width: 28,
                height: 28
              }
            }
          ]}
          cy={(cy: Cytoscape.Core) => cyActionHandlers(cy)}
        />
      </div>
    </div>
  );
};

export default memo(DependencyGraph);
