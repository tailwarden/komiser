import type { Meta, StoryObj } from '@storybook/react';
import Card from './Card';
import mockCardProps from './Card.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Card> = {
  title: 'Komiser/Card',
  component: Card,
  tags: ['autodocs'],
  argTypes: {
    icon: {}
  }
};

export default meta;
type Story = StoryObj<typeof Card>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockCardProps.base
  }
};
