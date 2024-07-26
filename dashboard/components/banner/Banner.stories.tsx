import type { Meta, StoryObj } from '@storybook/react';
import Banner from './Banner';
import mockBannerProps from './Banner.mocks';

const meta: Meta<typeof Banner> = {
  title: 'Komiser/Banner',
  component: Banner,
  tags: ['autodocs']
};

export default meta;
type Story = StoryObj<typeof Banner>;

export const Default: Story = {
  args: {
    ...mockBannerProps.base
  }
};

export const Primary: Story = {
  args: {
    ...mockBannerProps.primary
  }
};

export const Secondary: Story = {
  args: {
    ...mockBannerProps.secondary
  }
};
