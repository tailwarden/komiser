type VerticalSpacingProps = {
  size?: 'sm' | 'md' | 'lg';
};

/** Adds a top margin to separate components. It accepts a size prop which can take the values of 'sm', 'md' and 'lg', being 'md' the default. */
function VerticalSpacing({ size = 'md' }: VerticalSpacingProps) {
  return <div className={size === 'md' ? 'mt-6' : ''}></div>;
}

export default VerticalSpacing;
