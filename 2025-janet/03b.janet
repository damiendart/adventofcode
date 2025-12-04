# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split :s (group :values))
                                 :values (* (some (number :d)))}))

(defn largest-joltage [bank len]
  (var i (max ;(array/slice bank 0 (- (length bank) len -1))))
  (var idx (index-of i bank))
  (if (= len 1)
    (max ;bank)
    (+ (* i (math/pow 10 (- len 1))) (largest-joltage (array/slice bank (+ 1 idx)) (- len 1)))))

(defn solve [input]
  (reduce
    (fn [carry bank]
      (+ carry (largest-joltage bank 12)))
    0
    (peg/match input-parser input)))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
