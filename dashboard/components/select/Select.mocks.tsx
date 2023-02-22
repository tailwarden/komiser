import { SelectProps } from './Select';

const base: SelectProps = {
  label: 'Select',
  displayValues: ['Option 1', 'Option 2', 'Option 3'],
  onChange: () => {},
  values: ['o1', 'o2', 'o3']
};

const mockSelectProps = {
  base
};

export default mockSelectProps;
