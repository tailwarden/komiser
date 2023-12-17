import Image from 'next/image';
import Button from '../button/Button';

type Poses =
  | 'base'
  | 'boat'
  | 'celebrating'
  | 'confetti'
  | 'dam'
  | 'dumbbell'
  | 'fixing'
  | 'gift'
  | 'glasses'
  | 'greetings'
  | 'hat'
  | 'laptop'
  | 'laughing'
  | 'longsleeve'
  | 'reading'
  | 'rocket'
  | 'shirt'
  | 'whiteboard'
  | 'thinking'
  | 'working'
  | 'devops'
  | 'tablet';

export type EmptyStateProps = {
  title: string;
  message: string;
  action?: () => void | Element;
  secondaryAction?: () => void | Element;
  actionLabel?: string;
  secondaryActionLabel?: string;
  mascotPose?: Poses;
};

function EmptyState({
  title,
  message,
  action,
  secondaryAction,
  actionLabel = 'Guide to connect account',
  secondaryActionLabel = 'Report an issue',
  mascotPose
}: EmptyStateProps) {
  return (
    <div className="flex items-center justify-center">
      <div className="flex max-w-lg flex-col items-center justify-center gap-8 rounded-lg pt-8">
        {mascotPose && (
          <Image
            src={`/assets/img/purplin/${mascotPose}.svg`}
            width={150}
            height={190}
            alt="Purplin thinking"
            className="h-[190px] w-[150px]"
          />
        )}
        <div className="flex flex-col items-center gap-2 text-center">
          <p className="text-xl font-semibold text-gray-950">{title}</p>
          <p className="text-sm text-gray-700">{message}</p>
        </div>
        {action && (
          <div className="flex flex-wrap-reverse items-center gap-4 sm:flex-nowrap">
            {secondaryAction && (
              <Button size="lg" style="ghost" onClick={() => secondaryAction()}>
                {secondaryActionLabel}
              </Button>
            )}
            <Button size="lg" onClick={() => action()}>
              {actionLabel}
            </Button>
          </div>
        )}
      </div>
    </div>
  );
}

export default EmptyState;
