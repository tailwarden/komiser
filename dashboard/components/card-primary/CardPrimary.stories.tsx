import type { Meta, StoryObj } from '@storybook/react';
import CardPrimary from './CardPrimary';
import mockCardPrimaryProps from './CardPrimary.mock';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof CardPrimary> = {
  title: 'Komiser/Card Primary',
  component: CardPrimary,
  tags: ['autodocs'],
  argTypes: {
    type: {
      control: {
        type: 'inline-radio'
      }
    },
    showButton: {
      control: 'boolean'
    },
    showAvatar: {
      control: 'boolean'
    }
  }
};

export default meta;
type Story = StoryObj<typeof CardPrimary>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockCardPrimaryProps.base
  }
};
