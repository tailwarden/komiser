/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, useCallback, useLayoutEffect } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import Cytoscape from 'cytoscape';
// @ts-ignore
import COSEBilkent from 'cytoscape-cose-bilkent';

import { ReactFlowData } from './hooks/useDependencyGraph';

export type LayoutFlowProps = {
  data: ReactFlowData;
};

Cytoscape.use(COSEBilkent);
const LayoutFlow = ({ data }: LayoutFlowProps) => {
  const layout = { name: 'cose-bilkent', animate: true };
  let cy;

  return (
    <div className="relative h-full flex-1">
      <div className="h-[70vh]">
        <CytoscapeComponent
          className="h-full w-full"
          elements={CytoscapeComponent.normalizeElements({
            nodes: data.nodes,
            edges: data.edges
          })}
          layout={layout}
          stylesheet={[
            {
              selector: 'node',
              style: {
                width: 40,
                height: 40,
                content: 'data(label)',
                'background-image': node =>
                  node.data('resource') === 'AWS'
                    ? '/assets/img/dependency-graph/aws-node.svg'
                    : ''
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
        />
      </div>
    </div>
  );
};

export default LayoutFlow;
