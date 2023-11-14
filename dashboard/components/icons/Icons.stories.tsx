import type { Meta, StoryObj } from '@storybook/react';
import * as icons from '@components/icons';
import { SVGProps } from 'react';
import Tooltip from '@components/tooltip/Tooltip';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction

const IconsWrapper = (props: SVGProps<SVGSVGElement>) => (
  <div className="inline-flex w-full flex-wrap gap-2 p-2">
    {Object.entries(icons).map(([name, Icon]) => (
      <div key={name} className="relative">
        <div className="peer flex h-full flex-col items-center justify-center gap-2 rounded-md border bg-gray-200 p-3">
          <Icon {...props} />
          <p className="text-sm">{name}</p>
        </div>
        <Tooltip align="center">{`import { ${name} } from "@components/icons"`}</Tooltip>
      </div>
    ))}
  </div>
);

const meta: Meta<typeof IconsWrapper> = {
  title: 'Komiser/Icons',
  component: IconsWrapper,
  tags: ['autodocs'],
  argTypes: {}
};

export default meta;
type Story = StoryObj<typeof IconsWrapper>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    width: '24',
    height: '24'
  }
};
