import Cytoscape from 'cytoscape';

export const zoomLevelBreakpoint = 1.5;
export const maxZoom = 4;
export const minZoom = 0.25;
export const graphLayoutConfig = {
  name: 'cose-bilkent',
  // 'draft', 'default' or 'proof"
  // - 'draft' fast cooling rate
  // - 'default' moderate cooling rate
  // - "proof" slow cooling rate
  quality: 'proof',
  // Whether to include labels in node dimensions. Useful for avoiding label overlap
  nodeDimensionsIncludeLabels: true,
  // number of ticks per frame; higher is faster but more jerky
  refresh: 30,
  // Whether to fit the network view after when done
  fit: true,
  // Padding on fit
  padding: 5,
  // Whether to enable incremental mode
  randomize: true,
  // Node repulsion (non overlapping) multiplier
  nodeRepulsion: 4500,
  // Ideal (intra-graph) edge length
  idealEdgeLength: 75,
  // Divisor to compute edge forces
  edgeElasticity: 0.45,
  // Nesting factor (multiplier) to compute ideal edge length for inter-graph edges
  nestingFactor: 1,
  // Gravity force (constant)
  gravity: 0.25,
  // Maximum number of iterations to perform
  numIter: 2500,
  // Whether to tile disconnected nodes
  tile: true,
  // Type of layout animation. The option set is {'during', 'end', false}
  animate: 'end',
  // Duration for animate:end
  animationDuration: 500,
  // Amount of vertical space to put between degree zero nodes during tiling (can also be a function)
  tilingPaddingVertical: 100,
  // Amount of horizontal space to put between degree zero nodes during tiling (can also be a function)
  tilingPaddingHorizontal: 100,
  // Gravity range (constant) for compounds
  gravityRangeCompound: 1.5,
  // Gravity force (constant) for compounds
  gravityCompound: 1.0,
  // Gravity range (constant)
  gravityRange: 3.8,
  // Initial cooling factor for incremental layout
  initialEnergyOnIncremental: 0.5,
  nodeSeparation: 20000
};

export const nodeStyeConfig = {
  width(node) {
    return Math.max(2, Math.ceil(node.degree(false) / 2)) * 20;
  },
  height(node) {
    return Math.max(2, Math.ceil(node.degree(false) / 2)) * 20;
  },
  shape: 'ellipse',
  'text-opacity': 1,
  'font-size': 17,
  'background-color': 'white',
  'background-image': node => {
    switch (node.data('provider')) {
      case 'AWS':
        return '/assets/img/dependency-graph/aws-node.svg';
      case 'Civo':
        return '/assets/img/dependency-graph/civo-node.svg';
      default:
        return '';
    }
  },
  'background-height': 20,
  'background-width': 20,
  'border-color': '#EDEBEE',
  'border-width': 1,
  'border-style': 'solid',
  'transition-property': 'opacity',
  'transition-duration': 0.2,
  'transition-timing-function': 'linear'
} as Cytoscape.Css.Node;

export const edgeStyleConfig = {
  width: 1,
  'line-fill': 'linear-gradient',
  'line-gradient-stop-colors': ['#008484', '#33CCCC'],
  'line-style': edge => (edge.data('relation') === 'USES' ? 'solid' : 'dashed'),
  'curve-style': 'unbundled-bezier',
  'control-point-distances': edge => edge.data('controlPointDistances'),
  'control-point-weights': [0.25, 0.75]
} as Cytoscape.Css.Edge;

export const leafStyleConfig = {
  width: 28,
  height: 28,
  opacity: 1
} as Cytoscape.Css.Node;

export const edgeAnimationConfig = [
  {
    zoom: { level: 1 },
    easing: 'linear',
    style: {
      'line-dash-offset': 24,
      'line-dash-pattern': [4, 4]
    }
  },
  {
    duration: 4000
  }
];

export const nodeHTMLLabelConfig = {
  query: 'node', // cytoscape query selector
  halign: 'center', // title vertical position. Can be 'left',''center, 'right'
  valign: 'bottom', // title vertical position. Can be 'top',''center, 'bottom'
  halignBox: 'center', // title vertical position. Can be 'left',''center, 'right'
  valignBox: 'bottom', // title relative box vertical position. Can be 'top',''center, 'bottom'
  cssClass: 'dependency-graph-nodeLabel' // any classes will be as attribute of <div> container for every title
};
