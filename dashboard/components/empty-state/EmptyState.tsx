import { ReactElement, ReactNode } from 'react';
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
  | 'working';

export type EmptyStateProps = {
  title: string;
  message: string;
  action?: () => void | Element;
  actionLabel?: string;
  mascotPose?: Poses;
};

function EmptyState({
  title,
  message,
  action,
  actionLabel,
  mascotPose
}: EmptyStateProps) {
  return (
    <div className="flex h-[calc(100vh-156px)] items-center justify-center">
      <div className="flex flex-col items-center justify-center max-w-sm bg-white p-12 pb-0 rounded-lg">
        <p className="font-medium text-black-900">{title}</p>
        <div className="mt-2"></div>
        <p className="text-sm text-black-300 text-center">{message}</p>
        <div className="mt-8"></div>
        {action && (
          <>
            <Button size="lg" style="outline" onClick={() => action()}>
              {actionLabel}
            </Button>
            <div className="mt-8"></div>
          </>
        )}
        {mascotPose && (
          <picture className="h-[10rem] overflow-hidden">
            <img
              src={`/assets/img/purplin/${mascotPose}.svg`}
              className="w-48"
              alt="Purplin"
            />
          </picture>
        )}
      </div>
    </div>
  );
}

export default EmptyState;
