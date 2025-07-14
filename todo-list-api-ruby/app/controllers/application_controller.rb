class ApplicationController < ActionController::API
  after_action :set_jwt_token

  private

  def set_jwt_token
    if request.path == rodauth.logout_path
      response.headers["Authorization"] = nil
      return
    end

    if rodauth.use_jwt? && rodauth.valid_jwt?
      response.headers["Authorization"] = rodauth.session_jwt
    end
  end

  def current_user
    @current_user ||= rodauth.rails_account
  end
end
