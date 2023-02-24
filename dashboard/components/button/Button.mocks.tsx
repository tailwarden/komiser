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

const bulk: ButtonProps = {
  children: 'Bulk button',
  type: 'button',
  style: 'bulk',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const bulkOutline: ButtonProps = {
  children: 'Bulk button',
  type: 'button',
  style: 'bulk-outline',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const outline: ButtonProps = {
  children: 'Bulk button',
  type: 'button',
  style: 'outline',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const ghost: ButtonProps = {
  children: (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      fill="none"
      viewBox="0 0 24 24"
    >
      <path
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeMiterlimit="10"
        strokeWidth="2"
        d="M9.57 5.93L3.5 12l6.07 6.07M20.5 12H3.67"
      ></path>
    </svg>
  ),
  type: 'button',
  style: 'ghost',
  size: 'sm',
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

const deleteButtonGhost: ButtonProps = {
  children: 'Delete button ghost',
  type: 'button',
  style: 'delete-ghost',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {}
};

const mockButtonProps = {
  base,
  secondary,
  bulk,
  bulkOutline,
  outline,
  ghost,
  deleteButton,
  deleteButtonGhost
};

export default mockButtonProps;
