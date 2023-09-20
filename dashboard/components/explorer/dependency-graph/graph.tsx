/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, memo } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import Cytoscape, { EventObject } from 'cytoscape';
// @ts-ignore
import COSEBilkent from 'cytoscape-cose-bilkent';

import { ReactFlowData } from './hooks/useDependencyGraph';

export type LayoutFlowProps = {
  data: ReactFlowData;
};

const layout = { name: 'cose-bilkent', animate: true };

Cytoscape.use(COSEBilkent);
const LayoutFlow = ({ data }: LayoutFlowProps) => {
  const [initDone, setInitDone] = useState(false);

  const cyActionHandlers = (cy: Cytoscape.Core) => {
    if (!initDone) {
      cy.on('click', 'node', (evt: EventObject) => {
        console.info(`Clicked the node with ID ${evt.target.id()}`);
      });
      setInitDone(true);
    }
  };

  return (
    <div className="relative h-full flex-1">
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
                width: 40,
                height: 40,
                shape: 'ellipse',
                content: 'data(label)',
                'background-color': 'white',
                'background-image': node =>
                  node.data('resource') === 'AWS'
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
                'curve-style': 'bezier'
              }
            }
          ]}
          cy={(cy: Cytoscape.Core) => cyActionHandlers(cy)}
        />
      </div>
    </div>
  );
};

export default memo(LayoutFlow);
