import { ComponentStory, ComponentMeta } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import EmptyState, { EmptyStateProps } from './EmptyState';
import mockEmptyStateProps from './EmptyState.mocks';

export default {
  title: 'Components/Empty State',
  component: EmptyState,
  // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
  argTypes: {
    mascotPose: {
      control: { type: 'radio' }
    }
  }
} as ComponentMeta<typeof EmptyState>;

// More on component templates: https://storybook.js.org/docs/react/writing-stories/introduction#using-args
const Template: ComponentStory<typeof EmptyState> = args => (
  <EmptyState {...args} />
);

export const Default = Template.bind({});
// More on args: https://storybook.js.org/docs/react/writing-stories/args

Default.args = {
  ...mockEmptyStateProps.base,
  action: action('onClick')
} as EmptyStateProps;
