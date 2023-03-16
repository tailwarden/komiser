import Image from 'next/image';
import Button from '@/components/button/Button';

type InventoryViewAlertHasNoSlackAlertsProps = {
  createOrEditSlackAlert: (alertId?: number | undefined) => void;
};

function InventoryViewAlertHasNoSlackAlerts({
  createOrEditSlackAlert
}: InventoryViewAlertHasNoSlackAlertsProps) {
  return (
    <div className="rounded-lg bg-black-100 p-6">
      <div className="flex items-center gap-6">
        <Image
          src="/assets/img/purplin/tablet.svg"
          alt="Purplin"
          width={140}
          height={180}
          className="flex-shrink-0"
        />
        <div className="flex flex-col items-start gap-2 px-4">
          <p className="font-semibold text-black-900">
            It seems you have no alerts set up
          </p>
          <p className="text-sm text-black-400">
            Set up budget or resources alerts to stay updated on your cloud
            infrastructure
          </p>
          <span></span>
          <Button onClick={createOrEditSlackAlert}>Add alert</Button>
        </div>
      </div>
    </div>
  );
}

export default InventoryViewAlertHasNoSlackAlerts;
