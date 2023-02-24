import { SelectCheckboxProps } from './SelectCheckbox';

const base: SelectCheckboxProps = {
  label: 'Exclude',
  listOfResources: ['AWS', 'Kubernetes', 'Civo', 'Azure', 'Other'],
  exclude: [],
  setExclude: () => {}
};

const mockSelectCheckboxProps = {
  base
};

export default mockSelectCheckboxProps;
