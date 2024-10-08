import type { Meta, StoryObj } from '@storybook/react';
import mockCtaProps from './Cta.mock';
import Cta from './Cta';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Cta> = {
  title: 'Komiser/Cta',
  component: Cta,
  tags: ['autodocs'],
  argTypes: {
    action: {
      control: 'hidden'
    }
  }
};

export default meta;
type Story = StoryObj<typeof Cta>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockCtaProps.base
  }
};
