import type { Meta, StoryObj } from '@storybook/react';
import CardSkeleton from './CardSkeleton';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof CardSkeleton> = {
  title: 'Komiser/Card',
  component: CardSkeleton,
  tags: ['autodocs'],
  argTypes: {
    icon: {}
  }
};

export default meta;
type Story = StoryObj<typeof CardSkeleton>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Skeleton: Story = {
  args: {}
};
