import { InputProps } from './Input';

const base: InputProps = {
  id: 0,
  label: 'Text input',
  name: 'text',
  type: 'text',
  error: `Please provide a valid value.`,
  autofocus: true,
  action: () => {}
};

const mockButtonProps = {
  base
};

export default mockButtonProps;
