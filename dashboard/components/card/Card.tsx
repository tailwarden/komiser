import { ReactNode } from 'react';
import formatNumber from '../../utils/formatNumber';
import Tooltip from '../tooltip/Tooltip';

export type CardProps = {
  label: string;
  value: number;
  tooltip?: string;
  icon: ReactNode;
  formatter?: 'currency' | 'standard';
};

function Card({ label, value, tooltip, icon, formatter }: CardProps) {
  return (
    <div className="relative flex w-full items-center gap-4 rounded-lg bg-white px-6 py-8 text-gray-950 transition-colors">
      <div className="rounded-lg bg-gray-50 p-4" data-testid="icon">
        {icon}
      </div>
      <div className="peer flex flex-col">
        <p className="text-xl font-medium" data-testid="formattedNumber">
          {formatNumber(
            value,
            formatter === 'currency' ? 'currency' : undefined
          )}
        </p>
        <p className="text-sm text-gray-500">{label}</p>
      </div>
      {tooltip && <Tooltip>{tooltip}</Tooltip>}
    </div>
  );
}

export default Card;
