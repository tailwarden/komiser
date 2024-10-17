import type { Meta, StoryObj } from '@storybook/react';

import Modal, { ModalProps } from './Modal';

function ModalWrapper({
  children,
  isOpen,
  closeModal
}: Pick<ModalProps, 'children' | 'isOpen'> & { closeModal: () => void }) {
  return (
    <Modal isOpen={isOpen} closeModal={() => closeModal()}>
      {children}
    </Modal>
  );
}

const meta: Meta<typeof Modal> = {
  title: 'Komiser/Modal',
  component: ModalWrapper,
  decorators: [
    Story => <div style={{ margin: '3em', height: '200px' }}>{Story()}</div>
  ],
  tags: ['autodocs'],
  argTypes: {
    isOpen: {
      control: 'boolean'
    },
    children: {
      control: 'text'
    }
  }
};

export default meta;

type Story = StoryObj<typeof Modal>;

export const Primary: Story = {
  args: {
    children: 'Lorem Ipsum dolor'
  }
};
