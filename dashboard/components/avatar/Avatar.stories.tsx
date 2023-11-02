import type { Meta, StoryObj } from '@storybook/react';
import { allProviders, IntegrationProvider } from '@utils/providerHelper';
import Avatar from './Avatar';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Avatar> = {
  title: 'Komiser/Avatar',
  component: Avatar,
  tags: ['autodocs'],
  args: {
    size: 48
  }
};

export default meta;
type Story = StoryObj<typeof Avatar>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const AmazonWebServices: Story = {
  args: {
    avatarName: allProviders.AWS
  }
};

export const GoogleCloudPlatform: Story = {
  args: {
    avatarName: allProviders.GCP
  }
};

export const DigitalOcean: Story = {
  args: {
    avatarName: allProviders.DIGITAL_OCEAN
  }
};

export const Azure: Story = {
  args: {
    avatarName: allProviders.AZURE
  }
};

export const Civo: Story = {
  args: {
    avatarName: allProviders.CIVO
  }
};

export const Kubernetes: Story = {
  args: {
    avatarName: allProviders.KUBERNETES
  }
};

export const Linode: Story = {
  args: {
    avatarName: allProviders.LINODE
  }
};

export const Tencent: Story = {
  args: {
    avatarName: allProviders.TENCENT
  }
};

export const OCI: Story = {
  args: {
    avatarName: allProviders.OCI
  }
};

export const Scaleway: Story = {
  args: {
    avatarName: allProviders.SCALE_WAY
  }
};

export const MongoDBAtlas: Story = {
  name: 'MongoDB Atlas',
  args: {
    avatarName: allProviders.MONGODB_ATLAS
  }
};

export const Terraform: Story = {
  args: {
    avatarName: allProviders.TERRAFORM
  }
};

export const Pulumi: Story = {
  args: {
    avatarName: allProviders.PULUMI
  }
};

export const Slack: Story = {
  args: {
    avatarName: IntegrationProvider.SLACK
  }
};

export const Webhook: Story = {
  args: {
    avatarName: IntegrationProvider.WEBHOOK
  }
};
