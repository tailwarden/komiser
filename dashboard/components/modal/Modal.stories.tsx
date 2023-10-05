// import type { Meta, StoryObj } from '@storybook/react';
// import { useArgs, useState, useEffect } from 'preview-api';

// import Modal, { ModalProps } from './Modal';

// function ModalWrapper({
//   children,
//   isOpen
// }: Pick<ModalProps, 'children' | 'isOpen'>) {
//   const [_, updateArgs] = useArgs();
//   return (
//     <Modal isOpen={isOpen} closeModal={() => updateArgs({ isOpen: !isOpen })}>
//       {children}
//     </Modal>
//   );
// }

// const meta: Meta<typeof Modal> = {
//   title: 'Komiser/Modal',
//   component: ModalWrapper,
//   decorators: [
//     Story => (
//       <div style={{ margin: '3em' }}>
//         {/* 👇 Decorators in Storybook also accept a function. Replace <Story/> with Story() to enable it  */}
//         {Story()}
//       </div>
//     )
//   ],
//   tags: ['autodocs'],
//   argTypes: {
//     isOpen: {
//       control: 'boolean'
//     },
//     children: {
//       control: 'string'
//     }
//   }
// };

// export default meta;

// type Story = StoryObj<typeof Modal>;

// export const Primary: Story = {
//   args: {
//     children: 'Lorem Ipsum dolor'
//   }
// };
export default {};
