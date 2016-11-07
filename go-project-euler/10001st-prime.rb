def isPrime(number)
  isPrime = true
  counter = 0
  i = 1
  while i <= Math.sqrt(number)
    if number % i == 0
      counter = counter + 1
    end
    i = i + 1
    if counter >= 2
      isPrime = false
    end
  end
  return isPrime
end

def getPrimeAtPosition(upperNumber)
  counter = 0
  i = 0
  while counter <= upperNumber
    i = i + 1
    if isPrime(i)
      counter = counter + 1
    end
  end
  return i
end


puts getPrimeAtPosition(100001)
