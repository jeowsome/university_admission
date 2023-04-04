from hstest import StageTest, CheckResult, WrongAnswer, TestCase

input_1 = ["70", "90", "60"]
input_2 = ["50", "53", "78"]
input_3 = ["100", "84", "10"]
input_4 = ["50", "50", "50"]
input_5 = ["12", "6", "3"]
input_6 = ["70", "60", "50"]
threshold = 60.0
# positive_keywords = ["congratulation", "accepted"]
# negative_keywords = ["regret", "not", "offer you admission"]
positive_message = "Congratulations, you are accepted!"
negative_message = "We regret to inform you that we will not be able to offer you admission."


class TestAdmissionProcedure(StageTest):
    def generate(self):
        return [
            TestCase(stdin=input_1, attach=input_1),
            TestCase(stdin=input_2, attach=input_2),
            TestCase(stdin=input_3, attach=input_3),
            TestCase(stdin=input_4, attach=input_4),
            TestCase(stdin=input_5, attach=input_5),
            TestCase(stdin=input_6, attach=input_6),
        ]

    @staticmethod
    def get_mean(*numbers):
        numbers = [int(number) for number in numbers]
        return sum(numbers) / len(numbers)

    # @staticmethod
    # def check_keywords(correct_words, wrong_words, reply_line):
    #     correct_words = [word in reply_line for word in correct_words]
    #     wrong_words = [word in reply_line for word in wrong_words]
    #     return len(correct_words) == sum(correct_words) and not sum(wrong_words)

    @staticmethod
    def check_message(correct_message, wrong_message, reply_line):
        if correct_message in reply_line.capitalize():
            return True
        if wrong_message in reply_line.capitalize():
            return False

    @staticmethod
    def check_result(correct_mean, reply_line):
        # if correct_mean >= threshold:
        #     is_correct_output = TestAdmissionProcedure.check_keywords(positive_keywords, negative_keywords, reply_line)
        # else:
        #     is_correct_output = TestAdmissionProcedure.check_keywords(negative_keywords, positive_keywords, reply_line)

        if correct_mean >= threshold:
            is_correct_output = TestAdmissionProcedure.check_message(positive_message, negative_message, reply_line)
        else:
            is_correct_output = TestAdmissionProcedure.check_message(negative_message, positive_message, reply_line)

        if not is_correct_output:
            raise WrongAnswer("The second line of your output seem to contain a wrong message.\n"
                              "Make sure that you use the correct threshold \n"
                              "and output the message exactly as stated in the description.\n\n"
                              "If the mean score is >= 60.0 your program should output: \n\"{0}\"\n\n"
                              "If the mean score is < 60.0 your program should output: \n\"{1}\"".format(
                                positive_message, negative_message))

    def check(self, reply: str, attach: list):
        output = reply.lower().strip().split('\n')
        output = [line for line in output if line]
        if len(output) != 2:
            raise WrongAnswer("The output should contain 2 lines. \n"
                              "However, {0} lines were found in your output.".format(len(output)))
        correct_mean = round(self.get_mean(*attach), 2)
        try:
            output_mean = round(float(output[0].strip()), 2)
        except ValueError:
            raise WrongAnswer("The first line of your output is supposed to contain nothing but a number.\n")

        if output_mean != correct_mean:
            raise WrongAnswer("The mean score in your output is {0}.\n"
                              "However, the answer {1} was expected.".format(output_mean, correct_mean))

        self.check_result(correct_mean, output[1])
        return CheckResult.correct()


if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()
