class TasksController < ApplicationController
  before_action :authenticate_user!
  before_action :set_task, only: %i[show update destroy]

  # GET /tasks
  def index
    @tasks = Task.for_user(current_user)

    render json: {data: @tasks, meta: {total: @tasks.count}}, status: :ok
  end

  # GET /tasks/1
  def show
    render json: @task
  end

  # POST /tasks
  def create
    @task = current_user.tasks.new(task_params)

    if @task.save
      render json: @task, status: :created, location: @task
    else
      render json: @task.errors, status: :unprocessable_entity
    end
  rescue ActionController::ParameterMissing
    render json: {error: "Please provide all required parameters", required_parameters: %w[name priority status]}, status: :bad_request
  end

  # PATCH/PUT /tasks/1
  def update
    if @task.update(task_params)
      render json: @task
    else
      render json: @task.errors, status: :unprocessable_entity
    end
  end

  # DELETE /tasks/1
  def destroy
    @task.destroy!
  end

  private

  # Use callbacks to share common setup or constraints between actions.
  def set_task
    @task = Task.for_user(current_user).find(params.expect(:id))
  end

  # Only allow a list of trusted parameters through.
  def task_params
    params.expect(task: [:user_id, :name, :priority, :status])
  end

  def authenticate_user!
    rodauth.require_account
  end
end
