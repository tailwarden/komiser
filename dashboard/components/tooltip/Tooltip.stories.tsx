import type { Meta, StoryObj } from '@storybook/react';
import WarningIcon from '@components/icons/WarningIcon';
import Tooltip from './Tooltip';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Tooltip> = {
  title: 'Komiser/Tooltip',
  component: Tooltip,
  parameters: {
    docs: {
      description: {
        component:
          // eslint-disable-next-line no-multi-str
          'The tooltip component required to be wrapped inside an relative positioned container. There also need to be a trigger element which is required to have `className="peer"`!\
          In this example the strigger element is the Info icon.\
          The Storybook preview gives you an idea about possible parameters but might not work 100% because you should either define top **or** bottom, **not** both.\
          To allow to show all possible options, we define both top, bottom and left, right in this example. Please keep this in mind!'
      }
    }
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div className="flex h-96 items-center justify-center">
        <div className="relative h-[16px] w-[16px]">
          <WarningIcon className="peer" height="16" width="16" />
          <Story />
        </div>
      </div>
    )
  ],
  argTypes: {
    top: {
      control: {
        type: 'inline-radio'
      }
    },
    bottom: {
      control: {
        type: 'inline-radio'
      }
    },
    align: {
      control: {
        type: 'inline-radio'
      }
    },
    width: {
      control: {
        type: 'inline-radio'
      }
    }
  }
};

export default meta;
type Story = StoryObj<typeof Tooltip>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const TopTiny: Story = {
  args: {
    top: 'xs',
    align: 'left',
    width: 'lg',
    children: "That's a tooltip"
  }
};
