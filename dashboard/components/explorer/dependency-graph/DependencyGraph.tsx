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

const layout = {
  name: 'cose-bilkent',
  animate: 'end',
  nodeRepulsion: 10000,
  idealEdgeLength: 100
};

nodeHtmlLabel(Cytoscape.use(COSEBilkent));
const DependencyGraph = ({ data }: DependencyGraphProps) => {
  const [initDone, setInitDone] = useState(false);

  const loopAnimation = eles => {
    const ani = eles.animation(
      {
        zoom: { level: 1 },
        easing: 'linear',
        style: {
          'line-dash-offset': 24,
          'line-dash-pattern': [4, 4],
          'line-gradient-stop-positions': [0.1, 0.9]
        }
      },
      {
        duration: 4000
      }
    );

    ani
      .reverse()
      .play()
      .promise('complete')
      .then(() => loopAnimation(eles));
  };

  const cyActionHandlers = (cy: Cytoscape.Core) => {
    // make sure we did n ot init already, otherwise this will be bound more than once
    if (!initDone) {
      // Add click handler per node
      cy.on('click', 'node', (evt: EventObject) => {
        console.info(`Clicked the node with ID ${evt.target.id()}`);
      });
      // Add HTML labels for better flexibility
      // @ts-ignore
      cy.nodeHtmlLabel([
        {
          query: 'node', // cytoscape query selector
          halign: 'center', // title vertical position. Can be 'left',''center, 'right'
          valign: 'bottom', // title vertical position. Can be 'top',''center, 'bottom'
          halignBox: 'center', // title vertical position. Can be 'left',''center, 'right'
          valignBox: 'bottom', // title relative box vertical position. Can be 'top',''center, 'bottom'
          cssClass: 'dependency-graph-node-label', // any classes will be as attribute of <div> container for every title
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
                 }</p></div>`; // your html template here
          }
        }
      ]);
      // Add class to leave nodes so we can make them smaller
      cy.nodes().leaves().addClass('leaf');
      // Animate edges?
      cy.edges().forEach(loopAnimation);
      cy.on('zoom', event => {
        const currentZoomLevel = cy.zoom();
        const dim = 100 / currentZoomLevel;
        const edgeWidth = 10 / currentZoomLevel;
        cy.$('edge').css({
          opacity: currentZoomLevel <= 1.5 ? 0 : 1
        });
        cy.$('.leaf').css({
          opacity: currentZoomLevel <= 1.5 ? 0 : 1
        });

        Array.from(
          document.querySelectorAll('.dependency-graph-node-label'),
          e =>
            currentZoomLevel <= 1.5
              ? (e.style.opacity = 0)
              : (e.style.opacity = 1)
        );
      });
      // Make sure to tell we inited successfully and prevent another init
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
                'background-width': 20,
                'border-color': '#EDEBEE',
                'border-width': 1,
                'border-style': 'solid',
                'transition-property': 'opacity',
                'transition-duration': 0.2,
                'transition-timing-function': 'linear'
              }
            },
            {
              selector: 'edge',
              style: {
                width: 1,
                'line-fill': 'linear-gradient',
                'line-gradient-stop-colors': 'yellow green',
                'line-style': edge =>
                  edge.data('relation') === 'USES' ? 'solid' : 'dashed',
                'curve-style': 'unbundled-bezier',
                'control-point-distances': [2, -2],
                'control-point-weights': [0.15, 0.85]
              }
            },
            {
              selector: '.leaf',
              style: {
                width: 28,
                height: 28,
                opacity: 0
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
