import Image from 'next/image';
import React from 'react';
import Button from '@/components/button/Button';

function InventoryViewAlertHasNoSlackIntegration() {
  return (
    <div className="rounded-lg bg-black-100 p-6">
      <div className="flex items-center gap-4">
        <Image
          src="/assets/img/purplin/slack.svg"
          alt="Purplin"
          width={190}
          height={105}
          className="flex-shrink-0"
        />
        <div className="flex flex-col items-start gap-2">
          <p className="font-semibold text-black-900">
            To set up alerts, connect your Slack
          </p>
          <p className="text-sm text-black-400">
            By setting up the Slack integration, you&apos;ll be able to better
            track your cloud usage and costs, enabling you and your team to
            always be notified when it matters.
          </p>
          <a
            href="https://docs.komiser.io/docs/introduction/getting-started"
            target="_blank"
            rel="noreferrer"
            className="mt-2"
          >
            <Button>Learn how to connect Slack</Button>
          </a>
        </div>
      </div>
    </div>
  );
}

export default InventoryViewAlertHasNoSlackIntegration;
