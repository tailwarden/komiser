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

const outline: ButtonProps = {
  children: 'Outline button',
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
  size: 'xs',
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
  outline,
  ghost,
  deleteButton
};

export default mockButtonProps;
