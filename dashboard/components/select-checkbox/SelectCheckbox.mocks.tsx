import { SelectCheckboxProps } from './SelectCheckbox';

const base: SelectCheckboxProps = {
  label: 'Exclude',
  query: 'provider',
  exclude: [],
  setExclude: () => {}
};

const mockSelectCheckboxProps = {
  base
};

export default mockSelectCheckboxProps;
