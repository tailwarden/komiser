import type { Meta, StoryObj } from '@storybook/react';
import EmptyState from './EmptyState';
import mockEmptyStateProps from './EmptyState.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof EmptyState> = {
  title: 'Komiser/Empty State',
  component: EmptyState,
  tags: ['autodocs'],
  argTypes: {
    action: {}
  }
};

export default meta;
type Story = StoryObj<typeof EmptyState>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockEmptyStateProps.base
  }
};
