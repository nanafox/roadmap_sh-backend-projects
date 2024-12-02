class UnitConverterController < ApplicationController
  before_action :set_conversions

  VALID_CONVERSIONS = {
    weight: [
      [ "Kilograms", "kg" ],
      [ "Grams", "g" ],
      [ "Milligrams", "mg" ],
      [ "Pounds", "lb" ],
      [ "Ounces", "oz" ]
    ],
    length: [
      [ "Millimeters", "mm" ],
      [ "Centimeters", "cm" ],
      [ "Meters", "m" ],
      [ "Kilometers", "km" ],
      [ "Inches", "in" ],
      [ "Feet", "ft" ],
      [ "Yards", "yd" ],
      [ "Miles", "mi" ]
    ],
    temperature: [
      [ "Celsius", "C" ],
      [ "Kelvin", "K" ],
      [ "Fahrenheit", "F" ]
    ],
    volume: [
      [ "Milliliters", "ml" ],
      [ "Liters", "l" ],
      [ "Cubic Meters", "m3" ],
      [ "Cubic Centimeters", "cm3" ],
      [ "Gallons", "gal" ],
      [ "Quarts", "qt" ],
      [ "Pints", "pt" ],
      [ "Fluid Ounces", "fl oz" ]
    ]
  }

  def show
  end

  def create
    @type = params[:type].to_sym
    value = params[@type]
    unit_from = params[:unit_from]
    unit_to = params[:unit_to]

    @type_value = Float(value, exception: false)
    if @type_value.nil?
      flash[:alert] = "'#{value}' is an invalid input, use numbers only!"
      return render :show, status: :unprocessable_content
    end

    if VALID_CONVERSIONS.key?(@type)
      @results = perform_conversion(@type, @type_value, unit_from, unit_to)
      render :show, status: :ok
    else
      flash[:error] = "Invalid conversion type."
      render :show, status: :unprocessable_content
    end
  end

  private

  def set_conversions
    @conversions = VALID_CONVERSIONS
  end

  def perform_conversion(type, value, unit_from, unit_to)
    case type
    when :length
      length_conversion(value, unit_from, unit_to)
    when :weight
      weight_conversion(value, unit_from, unit_to)
    when :temperature
      temperature_conversion(value, unit_from, unit_to)
    when :volume
      volume_conversion(value, unit_from, unit_to)
    else
      "Conversion not supported for type #{type}."
    end
  end

  def length_conversion(value, unit_from, unit_to)
    conversion_factor = {
      "mm" => 0.001, "cm" => 0.01, "m" => 1, "km" => 1000,
      "in" => 0.0254, "ft" => 0.3048, "yd" => 0.9144, "mi" => 1609.34
    }
    convert_with_factor(value, unit_from, unit_to, conversion_factor)
  end

  def weight_conversion(value, unit_from, unit_to)
    conversion_factor = {
      "mg" => 0.000001, "g" => 0.001, "kg" => 1,
      "lb" => 0.453592, "oz" => 0.0283495
    }
    convert_with_factor(value, unit_from, unit_to, conversion_factor)
  end

  def temperature_conversion(value, unit_from, unit_to)
    case [ unit_from, unit_to ]
    when [ "C", "C" ] then value
    when [ "C", "K" ] then value + 273.15
    when [ "K", "C" ] then value - 273.15
    when [ "C", "F" ] then (value * 9 / 5) + 32
    when [ "F", "C" ] then (value - 32) * 5 / 9
    when [ "K", "F" ] then (value - 273.15) * 9 / 5 + 32
    when [ "F", "K" ] then (value - 32) * 5 / 9 + 273.15
    else "Unsupported temperature conversion: #{unit_from} to #{unit_to}"
    end
  end

  def volume_conversion(value, unit_from, unit_to)
    conversion_factor = {
      "ml" => 0.001, "l" => 1, "m3" => 1000, "cm3" => 0.001,
      "gal" => 3.78541, "qt" => 0.946353, "pt" => 0.473176, "fl oz" => 0.0295735
    }
    convert_with_factor(value, unit_from, unit_to, conversion_factor)
  end

  def convert_with_factor(value, unit_from, unit_to, factor)
    if factor[unit_from] && factor[unit_to]
      converted_value = value * (factor[unit_from] / factor[unit_to])
      Rails.logger.debug "#{value} * (#{factor[unit_from]} / #{factor[unit_to]})"
      "#{value} #{unit_from} = #{converted_value.round(2)} #{unit_to}"
    else
      "Invalid conversion units: #{unit_from} to #{unit_to}"
    end
  end
end
