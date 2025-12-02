# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def parser (peg/compile ~{
  :main (split "," (group :range))
  :range (sequence (number :d+) "-" (number :d+))
}))

(defn solve [input]
  (var sum 0)
  (map
    (fn [range]
      (for i (get range 0) (+ (get range 1) 1)
        (var n (string/format "%d" i))
        (if (even? (length n))
          (if (string/has-prefix? (string/slice n (/ (length n) 2)) n) (set sum (+ sum i)))
        )
      )
    )
    (peg/match parser input)
  )
  sum
)

(assert (= (solve "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124") 1227775554))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
