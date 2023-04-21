import EditIcon from '../icons/EditIcon';
import { ButtonProps } from './Button';

const base: ButtonProps = {
  children: 'Primary button',
  type: 'button',
  style: 'primary',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const secondary: ButtonProps = {
  children: 'Secondary button',
  type: 'button',
  style: 'secondary',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const ghost: ButtonProps = {
  children: 'Ghost button',
  type: 'button',
  style: 'ghost',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const text: ButtonProps = {
  children: 'Text button',
  type: 'button',
  style: 'text',
  disabled: false,
  loading: false,
  onClick: () => {}
};
const dropdown: ButtonProps = {
  children: (
    <>
      <EditIcon width={24} height={24} />
      Dropdown button
    </>
  ),
  type: 'button',
  style: 'dropdown',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const deleteButton: ButtonProps = {
  children: 'Delete button',
  type: 'button',
  style: 'delete',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const mockButtonProps = {
  base,
  secondary,
  ghost,
  text,
  dropdown,
  deleteButton
};

export default mockButtonProps;
