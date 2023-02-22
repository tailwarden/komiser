import { SelectProps } from './Select';

const base: SelectProps = {
  label: 'Select',
  values: ['o1', 'o2', 'o3'],
  displayValues: ['Option 1', 'Option 2', 'Option 3'],
  onChange: () => {}
};

const mockSelectProps = {
  base
};

export default mockSelectProps;
