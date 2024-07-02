(1...100).each do |i|
    if i % 3 == 0 && i % 5 == 0
      p "FizzBuzz"
    elsif i % 3 == 0
      p "Fizz"
    elsif i % 5 == 0
      p "Buzz"
    else
      p i.to_s
    end
end