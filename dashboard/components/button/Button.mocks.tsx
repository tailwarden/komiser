import EditIcon from '../icons/EditIcon';
import { ButtonProps } from './Button';

const base: ButtonProps = {
  children: 'Primary button',
  type: 'button',
  style: 'primary',
  size: 'lg',
  disabled: false,
  loading: false,
  onClick: () => {},
  href: '',
  target: ''
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

const linkButton: ButtonProps = {
  children: 'Link button',
  asLink: true,
  style: 'primary',
  size: 'lg',
  loading: false,
  href: 'https://komiser.io'
};

const newTabLinkButton: ButtonProps = {
  children: 'New Tab Link button',
  asLink: true,
  style: 'secondary',
  size: 'lg',
  loading: false,
  href: 'https://komiser.io',
  target: '_blank'
};

const mockButtonProps = {
  base,
  secondary,
  ghost,
  text,
  dropdown,
  deleteButton,
  linkButton,
  newTabLinkButton
};

export default mockButtonProps;
