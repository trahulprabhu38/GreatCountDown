import React from 'react';

const ProgressBar = ({ percentComplete }) => {
  const percentage = percentComplete || 0;

  return (
    <div className="progress-container">
      <div className="progress-bar">
        <div
          className="progress-fill"
          style={{ width: `${percentage}%` }}
        >
          <span className="progress-text">{percentage.toFixed(2)}%</span>
        </div>
      </div>
      <div className="progress-label">Progress</div>
    </div>
  );
};

export default ProgressBar;
