class Task < ApplicationRecord
  belongs_to :user

  validates :name, presence: true
  validates :status, presence: true

  enum :status, pending: 1, completed: 2, cancelled: 3
  enum :priority, low: 1, medium: 2, high: 3

  scope :for_user, ->(user) { where(user_id: user.id) }
end
