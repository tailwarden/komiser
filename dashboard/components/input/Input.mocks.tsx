import regex from '../../utils/regex';
import { InputProps } from './Input';

const base: InputProps = {
  id: 0,
  label: 'Text input',
  name: 'text',
  type: 'text',
  regex: regex.required,
  error: `Please provide a valid value.`,
  action: () => {}
};

const mockButtonProps = {
  base
};

export default mockButtonProps;
