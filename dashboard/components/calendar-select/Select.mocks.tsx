import { CalendarSelectProps } from './CalendarSelect';

const base: CalendarSelectProps = {
  value: 'o1',
  label: 'Select',
  values: ['o1', 'o2', 'o3'],
  displayValues: ['Option 1', 'Option 2', 'Option 3'],
  handleChange: () => {}
};

const mockSelectProps = {
  base
};

export default mockSelectProps;
