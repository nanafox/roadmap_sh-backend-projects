require "test_helper"

class TaskTest < ActiveSupport::TestCase
  def setup
    @user = users(:one)
  end

  test "task belongs to user" do
    task = @user.tasks.new(name: "Test Task", priority: "medium", status: "pending")
    assert_equal @user.id, task.user_id
  end

  test "task has a valid status" do
    task = @user.tasks.new(name: "Test Task", priority: "high", status: "pending")
    assert_includes Task.statuses, task.status
  end

  test "task has a valid priority" do
    task = @user.tasks.new(name: "Test Task", priority: "low", status: "pending")
    assert_includes Task.priorities, task.priority
  end

  test "task has default status and priority when not provided" do
    task = @user.tasks.new(name: "Test Task")
    assert_equal "pending", task.status
    assert_equal "low", task.priority
  end
end
