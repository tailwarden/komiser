import type { Meta, StoryObj } from '@storybook/react';
import Toast from './Toast';

const meta: Meta<typeof Toast> = {
  title: 'Komiser/Toast',
  component: Toast,
  tags: ['autodocs'],
  argTypes: {},
  decorators: [
    Story => (
      <div
        style={{
          minHeight: '130px',
          position: 'relative'
        }}
      >
        <Story />
      </div>
    )
  ]
};

export default meta;
type Story = StoryObj<typeof Toast>;

export const Primary: Story = {
  args: {
    title: 'Toast Title',
    message:
      'Street art same raclette freegan actually. Literally solarpunk disrupt bespoke af tousled hashtag meh hot chicken iPhone vegan fixie post-ironic quinoa.'
  }
};

export const WithError: Story = {
  args: {
    title: 'Toast Title',
    message:
      'Street art same raclette freegan actually. Literally solarpunk disrupt bespoke af tousled hashtag meh hot chicken iPhone vegan fixie post-ironic quinoa.',
    hasError: true
  }
};
