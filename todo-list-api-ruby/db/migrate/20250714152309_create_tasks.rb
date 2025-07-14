class CreateTasks < ActiveRecord::Migration[8.0]
  def change
    create_table :tasks do |t|
      t.belongs_to :user, null: false, foreign_key: true
      t.string :name, null: false
      t.integer :priority, null: false, default: 1  # low by default
      t.integer :status, null: false, default: 1 # pending by default

      t.timestamps
    end
  end
end
